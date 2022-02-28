<template>
  <div class="first-view">
    <my-header />
    <hero />
  </div>

  <categories-slider />
  <recent-articles :articles="recentArticles" />
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import MyHeader from "@/components/MyHeader.vue";
import Hero from "@/components/Hero.vue";
import CategoriesSlider from "@/components/CategoriesSlider.vue";
import RecentArticles from "@/components/RecentArticles.vue";

import RecentArticle from "@/types/RecentArticle";

export default defineComponent({
  name: "home",
  components: {
    MyHeader,
    Hero,
    CategoriesSlider,
    RecentArticles
  },
  setup() {
    const recentArticles = ref<RecentArticle[]>([
      {
        id: "1",
        title: "Coventry City Guide Including Coventry Hotels",
        time: 9,
        sharp: "Guide",
        description: "Some words"
      },
      {
        id: "2",
        title: "qwdqwdqwd",
        time: 10,
        sharp: "Guide",
        description: "Some worqwdqwddds"
      }
    ]);
    return {
      recentArticles
    };
  }

});
</script>
