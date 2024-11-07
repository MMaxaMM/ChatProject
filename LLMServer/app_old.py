import torch
from transformers import pipeline
from typing import Union
from fastapi import FastAPI    
from pydantic import BaseModel

app = FastAPI()

torch_device = 'cuda' if torch.cuda.is_available() else 'cpu'
print(f"##### Device is {torch_device} #####")

# Create model pipeline
pipe = pipeline(
    "text-generation",
    model="./data/models/Meta-Llama-3-8B-Instruct",
    model_kwargs={"torch_dtype": torch.bfloat16},
    device=torch_device,
)

print("##### Model is loaded #####")

DEFAULT_SYSTEM_PROMPT = "You are a helpful assistant called Llama-3. Write out your answer short and succinct!"
DEFAULT_MAX_TOKENS = 512

# Data model for making POST requests to /chat 
class ChatRequest(BaseModel):
    messages: list
    max_tokens: Union[int, None] = None


def generate(messages: list, max_tokens: int = None) -> str:

    max_tokens = max_tokens if max_tokens else DEFAULT_MAX_TOKENS

    prompt = pipe.tokenizer.apply_chat_template(
        messages,
        tokenize=False,
        add_generation_prompt=True
    )
        
    terminators = [
        pipe.tokenizer.eos_token_id,
        pipe.tokenizer.convert_tokens_to_ids("<|eot_id|>")
    ]

    outputs = pipe(
        prompt,
        max_new_tokens=max_tokens,
        eos_token_id=terminators,
    )

    generated_outputs = outputs[0]["generated_text"] # Full prompt
    text = generated_outputs[len(prompt):] # Just the response
    return text

def isSystemPrompt(msg):
    return True if msg["role"] == "system" else False

@app.post("/chat")
def chat(chat_request: ChatRequest):

    messages = chat_request.messages
    max_tokens = chat_request.max_tokens

    # Check system prompt, add one if necessary
    if not isSystemPrompt(messages[0]):
        msg = {"role": "system", "content": DEFAULT_SYSTEM_PROMPT}
        messages.insert(0, msg)

    print("##### Generating response... #####")

    response = generate(messages, max_tokens)
    return {"role": "assistent", "content": response}

if __name__ == "__main__":
    app.run(debug=False)