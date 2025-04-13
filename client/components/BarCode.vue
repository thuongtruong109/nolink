<script setup lang="ts">
import JsBarcode from "jsbarcode";
import { onMounted } from "vue";

defineProps<{
  text: string;
}>();

onMounted(() => {
  JsBarcode("#barcode", "1234", {
    lineColor: "#0aa",
    height: 40,
    displayValue: false,
  });
});
</script>

<template>
  <div>
    <img id="barcode" />
  </div>
</template>

<style scoped>
img {
  border-radius: 0.5rem;
  width: 10rem;
}
</style>
