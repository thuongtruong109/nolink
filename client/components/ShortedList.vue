<script setup lang="ts">
import ShortedItem from "./ShortedItem.vue";
import type { TShortedList } from "../types";

defineProps<{
  shortedList: TShortedList[];
}>();
</script>

<template>
  <ul>
    <ShortedItem
      :item="item"
      v-for="(item, i) in shortedList"
      :key="i"
      :isActive="i === 0"
    />
  </ul>
</template>

<style lang="scss" scoped>
ul {
  width: 100%;
  margin: 2rem 0;
}
</style>
