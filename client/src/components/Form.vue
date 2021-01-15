<template>
  <form v-on:submit="handleSubmit">
    <input v-model="urlInput" placeholder="Enter a URL to Shorten">
    <button type="submit">Go!</button>
  </form>
  <div v-if="shortened">
    <p>Woohoo!</p>
    <a :href="host + '/' + shortened">{{ shortened }}</a>
  </div>
  <div v-if="error">
    <p>{{ error }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { postURL, apiHost } from '../services/api';

type FormData = {
  urlInput: string;
  shortened: string | null;
  error: string | null;
  host: string;
};

export default defineComponent({
  name: 'Form',
  data() {
    return {
      urlInput: '',
      shortened: null,
      error: null,
      host: apiHost,
    } as FormData;
  },
  methods: {
    async handleSubmit(e: Event) {
      try {
        e.preventDefault();
        if (this.urlInput === '') {
          this.error = 'Field cannot be blank >:(';
          return;
        }
        const shortenedURL = await postURL(this.urlInput);
        this.shortened = shortenedURL.id;
      } catch (error) {
        this.error = error.message;
      }
    },
  },
});
</script>

<style>
a {
  color: #fff;
}
</style>
