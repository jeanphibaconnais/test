<template>
  <div id="app2">
    <p>App 2 - Structure en composant de l'app de todo list</p>
    <todo></todo>

  </div>

</template>

<script>
import Todo from './2/Todo';

export default {
  name: 'app2',
  components : { Todo }
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
