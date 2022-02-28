<template>
  <div class="recent-articles">
    <div class="recent-articles__container">
      <h2 class="recent-articles__title">Recent Articles</h2>
      <transition-group name="articles" tag="ul">
        <li class="recent-articles__item" v-for="(article, index) in articles" :key="index">
          <p>{{ article.title }}</p>
          <p>{{ article.description }}</p>
        </li>
      </transition-group>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import RecentArticle from "@/types/RecentArticle";

export default defineComponent({
  props: {
    articles: {
      required: true,
      type: Array as PropType<RecentArticle[]>
    }
  }
});
</script>

<style scoped lang="sass">
@import "../assets/sass/helps/mixins"
.recent-articles
  margin-bottom: 140px

  &__item
    margin-bottom: 2rem

  &__title
    @extend %heading-three

</style>