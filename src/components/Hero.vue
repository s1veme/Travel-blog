<template>
  <div class="hero hero__container">
    <hero-articles />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import HeroArticles from "@/components/HeroArticles.vue";

export default defineComponent({
  name: "Hero",
  components: {
    HeroArticles
  }
});
</script>

<style scoped lang="sass">
.hero
  height: calc(100% - 70px)
  width: 100%
  display: flex
  flex-direction: column
  align-items: flex-start
  justify-content: flex-end

  &__wrapper
    height: 100%
    width: 100%

  &__img
    height: 100%
    width: 100%
    object-fit: cover

</style>