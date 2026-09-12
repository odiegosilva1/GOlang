const API = "/tasks";
const list = document.getElementById("task-list");
const form = document.getElementById("form-task");
const input = document.getElementById("title");

async function loadTasks() {
    const res = await fetch(API);
    const tasks = await res.json();
    render(tasks);
}

function render(tasks) {
    list.innerHTML = "";
    tasks.forEach(task => {
        const li = document.createElement("li");
        if (task.done) li.classList.add("done");

        const checkbox = document.createElement("input");
        checkbox.type = "checkbox";
        checkbox.checked = task.done;
        checkbox.addEventListener("change", () => toggleTask(task));

        const span = document.createElement("span");
        span.className = "title";
        span.textContent = task.title;

        const del = document.createElement("button");
        del.className = "delete";
        del.textContent = "Excluir";
        del.addEventListener("click", () => deleteTask(task.id));

        li.append(checkbox, span, del);
        list.appendChild(li);
    });
}

async function createTask(title) {
    await fetch(API, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, done: false })
    });
    loadTasks();
}

async function toggleTask(task) {
    await fetch(`${API}/${task.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title: task.title, done: !task.done })
    });
    loadTasks();
}

async function deleteTask(id) {
    await fetch(`${API}/${id}`, { method: "DELETE" });
    loadTasks();
}

form.addEventListener("submit", e => {
    e.preventDefault();
    createTask(input.value.trim());
    input.value = "";
});

loadTasks();