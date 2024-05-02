
function populateForm() {
  let bringsWithValues = ['alcohol', 'food', 'partner', 'friend', 'dog', 'dye'];
  let bringsWith = userData.bringswith || [];
  let stays = userData.stays || [];

  bringsWith.forEach(function (value) {
    if (!bringsWithValues.includes(value)) {
      let input = document.createElement('input');
      let label = document.createElement('label');

      input.type = 'checkbox';
      input.id = value;
      input.name = 'bringswith';
      input.value = value;
      input.checked = true;

      label.htmlFor = value;
      label.appendChild(document.createTextNode(value));

      document.getElementById('newCheckboxes').appendChild(input);
      document.getElementById('newCheckboxes').appendChild(label);
      document.getElementById('newCheckboxes').appendChild(document.createElement('br'));
    } else {
      document.getElementById(value).checked = true;
    }
  });

  stays.forEach(function (value) {
    document.getElementById(value).checked = true;
  });
}

document.addEventListener('DOMContentLoaded', () => {
  populateForm()
});
