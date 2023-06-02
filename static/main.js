


document.addEventListener("DOMContentLoaded", function() {
  updateLists()
});


function inputTodo(button) {

  var todo_type = button.closest(".box").id;
  var inputId = "input_" + todo_type
  var inputText = document.getElementById(inputId).value;
    // Reset the input value
  

  fetch("/"+ todo_type ,{
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      name: inputText,
      done: false
    }),
  })
  .then(response => response.json())
  .then(data => displayList(data,todo_type))
  .catch(error => console.log(error));
  document.getElementById(inputId).value = '';
}
   
    
function displayList(data,todo_type) {

    var ulId = "list_" + todo_type;
    var listContainer = document.getElementById(ulId);
    listContainer.innerHTML = "";
  
    data.forEach(function(item) {
      var li = document.createElement("li");
      li.appendChild(document.createTextNode(item.name));
      if (item.done) {
        li.classList.add("done");
      }
      listContainer.appendChild(li);
  
      li.addEventListener("click", () => {
        li.classList.toggle("done");

        fetch("/toggle", {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            table: todo_type,
            name: item.name,
            done: item.done
          }),
        })
        .catch(error => console.log(error));
      });
    });
}

function deleteTodos(button) {
  var todo_type = button.closest(".box").id;
  fetch("/delete", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      table: todo_type
    }),
  })
  .catch(error => console.log(error));
  updateLists()
}


function updateLists() {
  var types = ["type_1", "type_2", "type_3", "type_4"];

  types.forEach(function(item) {
    fetch("/" + item)
      .then(response => response.json())
      .then(data => {
        displayList(data, item);
      })
      .catch(error => console.log(error));
  });
}
/*
function toggleTodo(params) {
  
}
*/