function initBringsWithAdder() {
  document.getElementById('add').addEventListener('click', (e) => {
    e.preventDefault();
    let input = document.getElementById('other');
    let newCheckbox = document.createElement('input');
    let label = document.createElement('label');
    let wrapper = document.createElement('div');
    let span = document.createElement('span');

    newCheckbox.type = 'checkbox';
    newCheckbox.id = input.value;
    newCheckbox.name = 'bringswith';
    newCheckbox.value = input.value;
    newCheckbox.checked = true;

    label.htmlFor = input.value;
    span.textContent = input.value;

    label.appendChild(newCheckbox);
    label.appendChild(span);
    wrapper.appendChild(label);
    wrapper.className = 'cat';

    document.getElementById('newCheckboxes').appendChild(wrapper);

    input.value = '';
  });

  document.getElementById('other').addEventListener('keypress', (e) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      document.getElementById('add').click();
    }
  });
}

document.addEventListener('DOMContentLoaded', () => {
  initBringsWithAdder()
});
