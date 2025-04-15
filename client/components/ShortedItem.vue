<script setup lang="ts">
import { computed, ref } from "vue";
import Check from "./icons/Check.vue";
import Copy from "./icons/Copy.vue";
import Arrow from "./icons/Arrow.vue";
import Code from "./Code.vue";
import type { TShortedList } from "../types";

const props = defineProps<{
  item: TShortedList;
  isActive: boolean;
}>();

// const config = useRuntimeConfig();
// const currentUrl = ref<string>(
//   computed(() => String(config.public.apiBase) + props.short) as unknown as string
// );

const currentUrl = ref<string>(
  computed(() => props.item.short) as unknown as string
);
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
      <span>Shorted:</span>
      <span v-if="isToggleMore">Origin:</span>
    </h4>
    <div class="item-content">
      <div class="link">
        <a :href="props.item.short" target="_blank">{{ props.item.short }}</a>
        <a :href="currentUrl" target="_blank" v-if="isToggleMore">{{
          props.item.origin
        }}</a>
      </div>
      <div v-if="isToggleMore"><Code :text="props.item.short" /></div>
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
        @click="onCopy(props.item.short)"
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
    overflow: hidden;

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
      padding: 0 0.5rem;

      a:nth-child(1) {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        max-width: 100%;
      }

      a:nth-child(2) {
        color: $black;
        font-weight: normal;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        max-width: 80%;
      }
    }
  }
}
</style>
