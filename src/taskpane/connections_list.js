let data = ["Ram", "Shyam", "Sita", "Gita"];
let list = document.getElementById("connection_list");
data.forEach((item) => {
let li = document.createElement("li");
li.setAttribute("role", "option")
li.innerText = item;

list.appendChild(li)});