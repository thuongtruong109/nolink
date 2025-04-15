<script setup lang="ts">
import { ref, reactive, watch, onMounted } from "vue";
import { saveAs } from "file-saver";
import { toString } from "qrcode";
import Download from "./icons/Download.vue";

const props = defineProps<{
  text: string;
}>();

const qr = reactive({
  output: "",
  color: ref("#000000"),
  color_timeout: null,
});

watch(
  () => qr.color,
  () => {
    if (qr.color_timeout) clearTimeout(qr.color_timeout);

    qr.color_timeout = setTimeout(() => {
      generateQr();
    }, 250);
  }
);

const generateQr = () => {
  if (!props.text) {
    qr.output = null;
    return;
  }
  toString(
    props.text,
    { color: { dark: qr.color, light: "#0000" } },
    (err: unknown, string: string) => {
      if (err) throw err;
      qr.output = string;
    }
  );
};

const downloadSvg = (e: Event) => {
  e.preventDefault();
  const blob = new Blob([qr.output], { type: "image/svg+xml" });
  saveAs(blob, props.text + ".svg");
};

const generateBarcode = () => {
  if (!props.text) {
    return;
  }
};

onMounted(() => {
  generateQr();
  generateBarcode();
});
</script>

<template>
  <div class="qrcode_response">
    <div class="left_response" v-html="qr.output" />

    <div class="right_response">
      <div class="color_track">
        <input v-model="qr.color" type="text" />
        <div>
          <input v-model="qr.color" type="color" />
        </div>
      </div>
      <button @click="downloadSvg" class="download_btn">
        <Download />
        <span>Download</span>
      </button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.qrcode_response {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  padding-top: 1rem;

  .left_response {
    background: white;
    border-radius: 0.5rem;
    width: 6rem;
    height: 6rem;
    padding: 0 !important;
  }

  .right_response {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    font-size: 1rem;
    margin-right: 1rem;

    .color_track {
      display: flex;
      align-items: center;
      background: white;
      padding: 0.25rem 0.5rem;
      border-radius: 0.5rem;
      border: 1px solid #ccc;
      margin: 0.5rem;

      input {
        background: transparent;
        width: 4.5rem;
      }
      div {
        margin-left: 0.5rem;
        border-radius: 50%;
        overflow: hidden;
        width: 1.25rem;
        height: 1.25rem;

        input {
          cursor: pointer;
          width: 200%;
          height: 200%;
          transform: translate(-25%, -25%);
        }
      }
    }

    .download_btn {
      @include button($black);
      width: fit-content;
    }
  }
}
</style>
