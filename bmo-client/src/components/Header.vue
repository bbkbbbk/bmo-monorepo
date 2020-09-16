<template>
    <div class="header">
        {{ label }}
    </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'CardForm',
  props: {
    label: {
      type: String,
      required: true,
    },
  },
});
</script>

<style scoped>
.header {
    color: white;
    background-color: #47D3B9;
    font-size: 30px;
    font-weight: bold;
    padding: 5px;
}
</style>
