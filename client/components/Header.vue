<template>
  <header>
    <div>
      <!-- https://img.icons8.com/external-vitaliy-gorbachev-flat-vitaly-gorbachev/58/external-link-social-media-vitaliy-gorbachev-flat-vitaly-gorbachev.png -->
      <img src="/favicon.svg" alt="logo" width="40" height="40" />
      <h1>Nolink</h1>
    </div>
    <blockquote>Ship your link in an easier way</blockquote>
  </header>
</template>

<style lang="scss" scoped>
header {
  padding: 1rem 0 2rem;

  div {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    font-size: 1.5rem;
    letter-spacing: 0.25rem;
    font-family: "Lobster Two", cursive;
    font-weight: 600;

    img {
      margin-right: 1rem;
      filter: drop-shadow(1px 1px 4px #9cbcff);
    }

    h1 {
      color: #80aaff;
      text-shadow: 1px 1px 4px #9cbcff;
    }
  }

  blockquote {
    color: gray;
    margin: 1rem 0 2rem;
    font-size: 1.2rem;
    text-align: center;
    font-family: "Dancing Script", cursive;
  }
}
</style>
