<template>
  <article class="article">
    <div class="article__wrapper">
      <div class="article__buttons">
        <button class="button__prev"> prev</button>
        <button class="button__prev"> next</button>
      </div>
      <p class="article__tag">Travels</p>
      <h1 class="article__title">When Is The Best Time
        of Year To Visit Japan?</h1>
      <p class="article__text">
        The good news for travelers is that there is no single best time of year to travel to Japan...
      </p>
      <router-link class="button__black button" to="/">Read more</router-link>
    </div>
  </article>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "HeroArticles"
})
</script>

<style scoped lang="sass">
@import "../assets/sass/helps/mixins"
.article
  background-color: #000
  max-width: 50%
  max-height: 100%

  &__buttons
    display: flex
    position: absolute
    right: -64px
    top: 0

  &__wrapper
    padding: 80px 50px 60px 50px
    position: relative

  &__tag,
  &__title
    color: #fff

  &__tag
    font-size: 1rem
    margin-bottom: 1.5rem

  &__title
    @extend %heading-one


  &__text
    color: #949799
    margin-bottom: 2.2rem
    @extend %subtitle-one


.button
  padding: 16px 34px
  font-weight: bold
  font-size: 1rem

  &__prev,
  &__next
    width: 64px
    height: 64px

  &__prev
    background-color: #fff

  &__next
    background-color: #000

</style>