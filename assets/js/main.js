function initBringsWithAdder() {
  document.getElementById('add').addEventListener('click', (e) => {
    e.preventDefault();
    let input = document.getElementById('other');
    let newCheckbox = document.createElement('input');
    let label = document.createElement('label');

    newCheckbox.type = 'checkbox';
    newCheckbox.id = input.value;
    newCheckbox.name = 'bringswith';
    newCheckbox.value = input.value;
    newCheckbox.checked = true;

    label.htmlFor = input.value;
    label.appendChild(document.createTextNode(input.value));

    document.getElementById('newCheckboxes').appendChild(newCheckbox);
    document.getElementById('newCheckboxes').appendChild(label);
    document.getElementById('newCheckboxes').appendChild(document.createElement('br'));

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

  let checkboxes = document.querySelectorAll('input[type="checkbox"]');
  checkboxes.forEach(checkbox => {
    checkbox.addEventListener('change', function () {
      let parent = this.closest('.checkbox-item--brings');
      if (this.checked) {
        parent.style.backgroundColor = '#0077b6';
      } else {
        parent.style.backgroundColor = '';
      }
    });
  });
});
