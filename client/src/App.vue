<template>
  <Header />
  <Form />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Form from './components/Form.vue';
import Header from './components/Header.vue';

export default defineComponent({
  name: 'App',
  components: {
    Form,
    Header,
  },
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin-top: 60px;
  color: #fff;
}
body {
  margin: 0;
  background-color: #34364d;
}
</style>
