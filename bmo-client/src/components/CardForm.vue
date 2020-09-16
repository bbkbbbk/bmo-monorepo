<template>
  <div class="input-box">
    <label class="block text-sm font-bold mb-1" for="term">
        Term
    </label>
    <input v-model="card.term" class="border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" type="text" placeholder="term">
    <label class="block text-sm font-bold mt-2 mb-1" for="defini
    tion">
        Definition
    </label>
    <input v-model="card.definition" class="border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" type="text" placeholder="definition">
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Card } from '@/type/index';

export default Vue.extend({
  name: 'CardForm',
  props: {
    card: {
      type: Card,
      required: true,
    },
  },
});
</script>

<style scoped>
.input-box {
    background-color: #D8F5F0;
    padding: 10px;
    margin: 20px 0;
}

input {
    border-color: white;
}
</style>
