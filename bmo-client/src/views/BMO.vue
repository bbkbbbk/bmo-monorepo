<template>
<div class="bmo-container">
  <div class="bmo-head">
      <span class="eye"></span>

          <div class="mouth"></div>

    <span class="eye"></span>

  </div>
  <span class="label">BMO</span>
</div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'BMO',
  components: {
  },
});
</script>

<style scoped>
.bmo-container {
    background-color: #47D3B9;
    height: 100vh;
    width: 100vw;
    display: flex;
    flex-direction: column;
}

.bmo-head {
    background-color: #cae3d0;
    height: 25%;
    width: 80%;
    max-width: 400px;
    margin: 200px auto 0 auto;
    display: flex;
    justify-content: center;
    align-items: center;
}

.eye {
    height: 25px;
    width: 25px;
    background-color: #1a1a1a;
    border-radius: 50%;
    display: inline-block;
    margin: 15%;
}

.mouth {
    border-bottom-left-radius: 110px;
    border-bottom-right-radius: 110px;
    border: 10px solid #1a1a1a;
    border-bottom: 0;
}

.label {
    font-size: 60px;
    font-weight: bold;
    color: #1a1a1a;
    margin: 20px;
}
</style>
