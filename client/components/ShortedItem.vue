<script setup lang="ts">
import { computed, ref } from "vue";
import Check from "./icons/Check.vue";
import Copy from "./icons/Copy.vue";
import Arrow from "./icons/Arrow.vue";
import QRCode from "./QRCode.vue";

const props = defineProps<{
  url: string;
  isActive: boolean;
}>();

// const config = useRuntimeConfig();
// const currentUrl = ref<string>(
//   computed(() => String(config.public.apiBase) + props.url) as unknown as string
// );

const currentUrl = ref<string>(computed(() => props.url) as unknown as string);
const isCopied = ref<boolean>(false);

const onCopy = (text: string) => {
  navigator.clipboard.writeText(text);
  isCopied.value = true;
  setTimeout(() => {
    isCopied.value = false;
  }, 1000);
};

const isToggleMore = ref<boolean>(props.isActive);
</script>

<template>
  <li :class="{ active: isToggleMore }">
    <h4>
      <span>Shorted: {{ props.isActive }}</span>
      <span v-if="isToggleMore">Origin:</span>
    </h4>
    <div class="item-content">
      <div class="link">
        <a :href="props.url" target="_blank">{{ props.url }}</a>
        <a :href="currentUrl" target="_blank" v-if="isToggleMore">{{
          currentUrl
        }}</a>
      </div>
      <div v-if="isToggleMore"><QRCode :text="url" /></div>
    </div>
    <div class="item-btn">
      <button
        class="more-btn"
        :class="{ active: isToggleMore }"
        @click="isToggleMore = !isToggleMore"
        type="button"
      >
        <Arrow />
      </button>
      <button
        @click="onCopy(props.url)"
        class="copy-btn"
        :disabled="isCopied"
        type="button"
      >
        <Copy v-if="!isCopied" />
        <span v-if="!isCopied">Copy</span>
        <Check v-else />
        <span v-if="isCopied">Copied</span>
      </button>
    </div>
  </li>
</template>

<style lang="scss" scoped>
li {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin: 0.5rem 0;
  background: $light-blue;
  padding: 0.5rem;
  border-radius: 0.5rem;
  width: 100%;
  font-weight: bold;

  &.active {
    border: 1px solid $blue;
  }

  h4 {
    span {
      padding: 0.5rem 0;
      display: block;
    }
  }

  .item-btn {
    @include flex-center;
    font-size: 0.8rem;

    button {
      &.more-btn {
        @include button($semi-blue);
        color: #6e6e6e;
        box-shadow: none;

        svg {
          transform: rotate(90deg);
          transition: all 0.2s ease-in-out;
        }
        &.active {
          svg {
            transform: rotate(0deg);
          }
        }
      }
      &.copy-btn {
        @include button($blue);
      }
    }
  }

  .item-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    a {
      @include link($blue);
      padding: 0.5rem 0;
    }

    .link {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      width: 100%;
      text-align: center;

      a:nth-child(2) {
        color: $black;
        font-weight: normal;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
    }
  }
}
</style>
