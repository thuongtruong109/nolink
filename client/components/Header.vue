<template>
  <header>
    <div>
      <img src="/favicon.svg" alt="logo" />
      <h1>nolink</h1>
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

    img {
      width: 2.5rem;
      height: 2.5rem;
      object-fit: cover;
      object-position: center;
      margin-right: 0.5rem;
    }

    h1 {
      color: $yellow;
      text-shadow: 1px 1px 8px rgba(255, 205, 2, 0.4);
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
