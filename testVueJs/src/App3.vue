<template>
  <div id="app3">
    <h2>App 3</h2>
      <todo-item v-for="item in groceryList"
        v-bind:todo="item"
        v-bind:key="item.id">
      </todo-item> 
  </div>
</template>

<script>
export default {
  name: 'app3',
  data() {
    return{
      todos : [
        { text : 'Test d\'array' },
        { text : 'Second occurs' },
        { text : 'Third occurs' }
      ],
      groceryList: [
        { id: 0, text: 'Légumes' },
        { id: 1, text: 'Fromage' },
        { id: 2, text: 'Tout ce que les humains sont supposés manger' }
      ]
    }
  },
}
</script>