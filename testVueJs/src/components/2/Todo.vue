<template>
    <div id="todo">
        <form v-on:submit.prevent="addTask(newTask)">
            <p><input v-model="newTask"> <button>Ajouter</button></p>
        </form>

        <ul>
            <li v-for="task in tasks">{{task}}</li>
        </ul> 
    </div>
    
</template>

<script>
export default {
    name: "todo",
	data: function() {
		return {
			task : {},
            tasks : [],
            newTask : null
		};
	},
	methods: {
        addTask: function(task){
            this.tasks.push(task); 
            this.newTask = ""; 
        }
    }
}
</script>

