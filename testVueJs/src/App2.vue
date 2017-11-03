<template>
  <div id="app2">
    <p v-if="show"> {{ message }} </p>
    <input v-model="message"/>

  </div>
</template>


<script>
export default {
  name: 'app2',
  data() { 
      return {
        message : 'test ici',
        show : true
      }
  }
}
</script>