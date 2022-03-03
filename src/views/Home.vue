<template>
  <div class="first-view">
    <my-header />
    <hero />
  </div>
  <categories-slider />
  <div class="content content__container">
    <div class="content__wrapper">
      <recent-articles :articles="recentArticles" />
      <who />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import MyHeader from "@/components/MyHeader.vue";
import Hero from "@/components/Hero.vue";
import CategoriesSlider from "@/components/CategoriesSlider.vue";
import RecentArticles from "@/components/RecentArticles.vue";
import RecentArticle from "@/types/RecentArticle";
import Who from "@/components/Who.vue";


export default defineComponent({
  name: "home",
  components: {
    MyHeader,
    Hero,
    CategoriesSlider,
    RecentArticles,
    Who
  },
  setup() {
    const recentArticles = ref<RecentArticle[]>([
      {
        id: "1",
        title: "Coventry City Guide Including Coventry Hotels",
        time: 9,
        sharp: "#Guide",
        description:
          "The diseases most commonly seen in travellers are diarrhoea, malaria (if you travel in a malaria-infested area), accidents (when travelling by automobile or swimming), wound infections and…",
        image: require("../assets/articles/1.jpg")
      },
      {
        id: "2",
        title: "Planning a Trip to Japan in the Time of Covid",
        time: 10,
        sharp: "#Guide",
        description: "Whether you’re just daydreaming for the moment, or planning a trip for the future, we’ve compiled essential information about visiting Japan in the time of Covid-19...",
        image: require("../assets/articles/2.jpg")
      }
    ]);


    return {
      recentArticles
    };
  }

});
</script>
<style lang="sass">
.content
  width: 100%

  &__wrapper
    width: 100%
    display: flex
</style>
