<script setup lang="ts">
import { ref } from "vue";
import Footer from "./components/Footer.vue";
import Header from "./components/Header.vue";
import ShortedList from "./components/ShortedList.vue";
import Cut from "./components/icons/Cut.vue";
import { EBUTTON_OPTION } from "./enums/option";
import { FetchMethod } from "./services";
import type { TShortedList } from "./types";

const config = useRuntimeConfig();

const url = ref<string>("");

const toolOption = ref<string>(EBUTTON_OPTION.SHORTEN);

const shortedList = ref<TShortedList[]>([]);

const onClearInput = () => {
  url.value = "";
  toolOption.value = EBUTTON_OPTION.SHORTEN;
};

// https://nuxt.com/docs/api/composables/use-cookie#usecookie

// onMounted(() => {
//   getAllShorted();
// });

// const getAllShorted = async () => {
//   const ipData = await FetchMethod(config.public.ipSource,
//     {
//       method: "GET",
//       mode: 'no-cors',
//     }
//   );
//   const allData = await FetchMethod(
//     `${config.public.apiBase}/own?ip=${ipData?.ip}`,
//     {
//       method: "POST",
//       mode: 'no-cors',
//     }
//   );

//   shortedList.value = allData;
// };

const onShorten = async () => {
  if (!url.value) return;
  toolOption.value = EBUTTON_OPTION.SHORTEN;

  const shortenData = await FetchMethod(`${config.public.apiBase}/`, {
    method: "POST",
    body: JSON.stringify({ url: url.value }),
  });

  shortedList.value.unshift({
    short: shortenData.short,
    origin: url.value,
    createdAt: new Date().toLocaleString("en-US", {
      timeZone: "UTC",
    }),
  });
  onClearInput();
};
</script>

<template>
  <main>
    <Header />

    <section>
      <div class="form">
        <input
          type="search"
          name="url"
          id="url"
          placeholder="Enter URL"
          v-model="url"
          required
        />

        <button
          :disabled="!url"
          alt="Shorten"
          title="shorten_btn"
          class="short_btn"
          @click="onShorten"
        >
          <Cut />
          <span>Shorten</span>
        </button>
      </div>

      <ShortedList :shortedList="shortedList" />
    </section>

    <Footer />
  </main>
</template>

<style lang="scss">
main {
  font-family: "Open Sans", sans-serif;
  font-size: 14px;
  color: #333;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 100vh;
  overflow: hidden;
  background-image: url("/landing.jpg");
  background-repeat: no-repeat;
  background-size: cover;
  background-position: 100% 100%;
  background-attachment: fixed;

  section {
    border-radius: 0.5rem;
    @include flex-center;
    flex-direction: column;

    .form {
      display: flex;
      align-items: center;
      justify-content: space-between;
      border-radius: 0.5rem;
      width: 100%;

      input[type="search"] {
        font-size: 0.9rem;
        border: none;
        box-shadow: 0px 5px 10px 0px rgba(0, 0, 0, 0.1);
        padding: 0.6rem 1rem;
        width: 100%;
      }

      .short_btn {
        @include button($red);
        padding: 1.1rem 0.8rem;
      }
    }
  }
}

@media (min-width: 400px) {
  section {
    width: 320px;
  }
}

@media (min-width: 500px) {
  section {
    width: 420px;
  }
}

@media (min-width: 640px) {
  section {
    width: 500px;
  }
}

@media (min-width: 768px) {
  section {
    width: 640px;
  }
}

@media (min-width: 850px) {
  section {
    width: 768px;
  }
}

@media (min-width: 1280px) {
  section {
    width: 850px;
  }
}
</style>
