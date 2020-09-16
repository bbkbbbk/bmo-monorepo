<template>
<div class="create-set-box">
  <Header label="CREATE SET"/>
  <div class="form-container">
    <div class="card-container">
        <div class="flex items-center border-b border-teal-500 py-2">
            <input v-model="title" class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none" type="text" placeholder="Title">
        </div>
        <div v-for="(card, index) in this.cards" :key="index">
            <CardForm :card="card" />
        </div>
    </div>
        <button class="addBtn py-2 px-4 rounded-full" @click="addCard">+</button>
  </div>
  <div class="footer">
        <button class="createBtn py-2 px-4 rounded" @click="handleCreate" :disabled="isDisabled" >Create</button>
  </div>
</div>
</template>

<script lang="ts">
import Vue from 'vue';
import Header from '@/components/Header.vue';
import CardForm from '@/components/CardForm.vue';
import { Card, Cards } from '@/type/index';

export default Vue.extend({
  name: 'CreateSet',
  components: {
    CardForm,
    Header,
  },
  data() {
    return {
      title: '' as string,
      cards: [] as Cards,
    };
  },
  created(): void {
    this.addCard();
  },
  methods: {
    handleCreate(): void {
      console.log(this.title);
      console.log(this.cards);
    },
    addCard(): void {
      const card = new Card('', '');
      this.cards.push(card);
    },
  },
  computed: {
    isDisabled(): boolean {
      const isTitleEmpty = this.title === '';
      const isCardEmpty = this.cards[0].term === '';
      return isTitleEmpty || isCardEmpty;
    },
  },
});
</script>

<style scoped>
.form-container {
    display: block;
}

.card-container {
    max-width: 500px;
    margin: auto;
    padding: 15px;
    text-align: left;
    color: #47D3B9;
}

.addBtn {
    background-color: #47D3B9;
    color: white;
    font-size: 20px;
    font-weight: bold;
}

.createBtn {
    background-color: transparent;
    color: #47D3B9;
    border: 1px solid #47D3B9;
    margin: 2em 1em;
}

.createBtn:active:hover {
    background-color: #47D3B9;
    color: white;
    border: #47D3B9;
}

.createBtn:disabled {
  color: #d1d1d1;
  border-color: #d1d1d1;
}

.footer {
    text-align: end;
}
</style>
