<template>
  <div id="app3">
    <p>App 3 - Structure en composant de l'app de todo list le paramètre</p>
    
    <input v-model="message" />

    <todo v-bind:message="message"></todo>

  </div>

</template>

<script>
import Todo from './3/Todo';

export default {
  name: 'app3',
  components : { Todo },
  data: function () {
    return { 
      message : null
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

  color: #2c3e50;
  margin-top: 0px;
}
</style>
