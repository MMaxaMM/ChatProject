const div = document.getElementById('runLine')

const circle = '<span id="circle" class="inline mycircle"> ●</span>'

div.innerHTML = circle

const lines = [['Ка', 'к ', 'сд', 'ать', ' г', 'ра', 'фы', ' Ми', 'ха', 'илу', ' Анд', 'рее', 'вичу', '?'],
               ['Ко', 'гда', ' Дми', 'трий', ' Иго', 'рев', 'ич', ' п', 'рий', 'дёт', ' на', ' па', 'ру', '?'],
               ['Что', ' та', 'кое', ' си', 'нта', 'ксич', 'еск', 'ая ', 'омо', 'ним', 'ия', '?'],
               ['Как ', 'сде', 'ла', 'ть ',  'НИРС', ' за', ' два ', 'ча', 'са ', 'до ', 'сда', 'чи?'],
               ['Как ', 'под', 'нять ', 'лок', 'аль', 'ный ', 'сер', 'вер ', 'в ', 'Cou', 'nter', '-Str', 'ike?'],
               ['Как ', 'пра', 'вил', 'ьно ', 'ста', 'вить ', 'уда', 'рен', 'ие: ', 'обес', 'пече', 'ние ', 'или ', 'обес', 'пече', 'ние?'],
               ['Сущ', 'ест', 'вует', ' ли ', 'Лен', 'инск', 'ая ', 'ком', 'ната ', 'и ', 'как ', 'её ', 'най', 'ти?']]

let i = 0 // Line
let j = 0 // Token

function runLine(){
    if (j == lines[i].length){
        j = 0
        div.innerHTML = circle
        if(i == lines.length-1){
            i = 0
        }
        else{
            i++
        }
    }
    else{
        const token = lines[i][j]
        const span = `<span class="inline">${token}</span>`
        document.getElementById('circle').remove()
        div.innerHTML += span
        div.innerHTML += circle
        j++
    }

    if (j == lines[i].length){
        done = setTimeout('runLine()', 3000)
    }
    else{
        done = setTimeout('runLine()', 120)
    }
}

runLine()