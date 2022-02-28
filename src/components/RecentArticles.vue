<template>
  <div class="recent-articles">
    <div class="recent-articles">
      <h2 class="heading-three">Recent Articles</h2>
      <transition-group name="articles" tag="ul">
        <li class="recent-articles__item" v-for="(article, index) in articles" :key="index">
          <div class="recent-articles__item-img">
            <img :src="article.image" :alt="article.title">
          </div>
          <div class="recent-articles__item-info">
            <div class="recent-articles__item-top">
              <p class="paragraph">
                {{ article.sharp }}
                <span class="recent-articles__item-top-span">
                {{ article.time }} minutes reading
                </span>
              </p>
            </div>
            <p class="recent-articles__item-title heading-two">
              {{ article.title }}
            </p>
            <p class="recent-articles__item-description subtitle-two">
              {{ article.description }}
            </p>
            <router-link to="/" class="button button__black">Read more</router-link>
          </div>
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
.button
  padding: 16px 34px
  font-weight: bold

.recent-articles
  margin-bottom: 140px

  &__item
    margin-bottom: 2rem
    display: flex
    align-items: center
    max-width: 80%

    &-title
      margin-bottom: 8px

    &-img
      width: 460px
      height: 400px
      background-color: #6f6f6f
      margin-right: 30px

      img
        height: 100%
        width: 100%
        object-fit: cover

    &-info
      width: 100%

    &-top-span
      color: #949799

    &-description
      margin-bottom: 24px
      color: #949799


.articles
  transition: all 0.5s ease-out
</style>