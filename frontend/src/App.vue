<template>
  <Navbar />
  <div class="mt-3 content-wrapper">
    <router-view v-if="$store.state.auth.checked"/>
    <div class="h-100" style="text-align:center" v-else>
      <Spinner size="8rem" thickness="1.25rem"/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

import Navbar from '@/components/navbar';
import Spinner from '@/components/misc/Spinner.vue';
import Center from '@/components/misc/Center.vue';

export default defineComponent({
  name: 'App',
  components: {
    Navbar,
    Spinner,
    Center
  }
})
</script>

<style>
  html, body, #app {
    height: 100%;
  }

  #app {
    display: flex;
    flex-direction: column;
  }

  .content-wrapper {
    flex: 1;
  }
</style>