<template>
  <div id="app4">
    <p>App 4 - Structure en composant de l'app de todo list avec communication enfant => parent</p>
    
    <todo v-bind:message="message" @addElement="addElementMessage"></todo>

    <p>
      dernier élément ajouté : {{ lastMessage }}
    </p>

  </div>

</template>

<script>
import Todo from './4/Todo';

export default {
  name: 'app4',
  components : { Todo },
  data: function () {
    return { 
      message : null,
      lastMessage : null
    }
  },
  methods : {
    addElementMessage (message) {
      this.lastMessage = message;
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
