<template>
  <div class="hero hero__container">
    <hero-articles :articles="heroArticles"/>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import HeroArticles from "@/components/HeroArticles.vue";
import HeroArticle from "@/types/HeroArticle";

export default defineComponent({
  name: "Hero",
  components: {
    HeroArticles
  },
  setup() {
    const heroArticles = ref<HeroArticle[]>([
      {
        sharp: "Travels",
        title: "When Is The Best Time of Year To Visit Japan?",
        text: "The good news for travelers is that there is no single best time of year to travel to Japan..."
      },
      {
        sharp: "Traveqwdqwls",
        title: "qwdqwdqwdqw?",
        text: "The qwdqwd dddwqd for travelers is that there is no single best time of year to travel to Japan..."
      }
    ]);
    return {
      heroArticles
    };
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