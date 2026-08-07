const categories = [
  {name:'Завтраки', ico:'🍳', key:'breakfast'},
  {name:'Супы', ico:'🍲', key:'soup'},
  {name:'Салаты', ico:'🥗', key:'salad'},
  {name:'Основные блюда', ico:'🍽️', key:'main'},
  {name:'Десерты', ico:'🍰', key:'dessert'},
  {name:'Выпечка', ico:'🥐', key:'baking'},
];

const magnetsEl = document.getElementById('magnets');
const selected = new Set();

categories.forEach((cat, i) => {
  const label = document.createElement('label');
  label.className = 'magnet';
  label.style.animationDelay = (i * 45) + 'ms';
  label.style.transform = 'rotate(' + (((i % 5) - 2) * 1.6) + 'deg)';
  label.innerHTML =
    '<input type="checkbox" class="magnet-input" value="' + cat.key + '">' +
    '<span class="ico">' + cat.ico + '</span>' + cat.name;
  const input = label.querySelector('input');
  input.addEventListener('change', () => {
    if (input.checked) { selected.add(cat.key); label.classList.add('on'); }
    else { selected.delete(cat.key); label.classList.remove('on'); }
    document.getElementById('hint').textContent = selected.size
      ? 'Выбрано категорий: ' + selected.size
      : 'Отметь хотя бы одну категорию блюд';
  });
  magnetsEl.appendChild(label);
});

const timeRange = document.getElementById('timeRange');
const timeVal = document.getElementById('timeVal');
timeRange.addEventListener('input', () => { timeVal.textContent = timeRange.value + ' мин'; });

const calRange = document.getElementById('calRange');
const calVal = document.getElementById('calVal');
calRange.addEventListener('input', () => { calVal.textContent = calRange.value + ' ккал'; });

const ingInput = document.getElementById('ingInput');
const ingAddBtn = document.getElementById('ingAddBtn');
const ingTags = document.getElementById('ingTags');
const addedIngredients = [];

function addIngredient(){
  const val = ingInput.value.trim();
  if (!val) return;
  addedIngredients.push(val);
  const tag = document.createElement('span');
  tag.className = 'ing-tag';
  tag.textContent = val;
  const rm = document.createElement('button');
  rm.type = 'button'; rm.textContent = '×'; rm.setAttribute('aria-label', 'Убрать ' + val);
  rm.addEventListener('click', () => { tag.remove(); const i = addedIngredients.indexOf(val); if (i > -1) addedIngredients.splice(i, 1); });
  tag.appendChild(rm);
  ingTags.appendChild(tag);
  ingInput.value = '';
  ingInput.focus();
}
ingAddBtn.addEventListener('click', addIngredient);
ingInput.addEventListener('keydown', (e) => { if (e.key === 'Enter') { e.preventDefault(); addIngredient(); } });

const resultEl = document.getElementById('result');
const loadingDots = document.getElementById('loadingDots');
const resName = document.getElementById('resName');
const resTime = document.getElementById('resTime');
const resCal = document.getElementById('resCal');
const resIngredients = document.getElementById('resIngredients');
const resSteps = document.getElementById('resSteps');

document.getElementById('goBtn').addEventListener('click', () => {
  const hint = document.getElementById('hint');
  if (selected.size === 0) {
    hint.textContent = 'Сначала отметь хотя бы одну категорию';
    return;
  }
  hint.textContent = '';

  resultEl.classList.add('show');
  loadingDots.style.display = 'flex';
  resName.textContent = '';
  resTime.textContent = '';
  resCal.textContent = '';
  resIngredients.innerHTML = '';
  resSteps.innerHTML = '';

  const payload = {
    categories: Array.from(selected),
    timeMinutes: Number(timeRange.value),
    calories: Number(calRange.value),
    ingredients: addedIngredients,
  };

  console.log(payload);

  fetch('/ask', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  })
    .then(res => res.json())
    .then(data => {
      loadingDots.style.display = 'none';
      resName.textContent = data.name;
      resTime.textContent = data.timeMinutes + ' мин';
      resCal.textContent = data.calories + ' ккал';
      data.ingredients.forEach(item => {
        const li = document.createElement('li');
        li.textContent = item;
        resIngredients.appendChild(li);
      });
      data.steps.forEach(step => {
        const li = document.createElement('li');
        li.textContent = step;
        resSteps.appendChild(li);
      });
    })
    .catch(() => {
      loadingDots.style.display = 'none';
      resName.textContent = 'Не удалось получить ответ. Попробуй ещё раз.';
    })
});