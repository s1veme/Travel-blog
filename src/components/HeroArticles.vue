<template>
  <article class="article">
    <div class="hero-slider">
      <swiper
        :slidesPerView="1"
        :navigation="true"
        class="hero-swiper"
      >
        <swiper-slide
          class="article__wrapper"
          v-for="(article, index) in articles"
          :key="index">
          <p class="article__tag">{{ article.sharp }}</p>
          <h1 class="article__title">{{ article.title }}</h1>
          <p class="article__text"> {{ article.text }}</p>
          <router-link class="button__black button" to="/">Read more</router-link>
        </swiper-slide>
      </swiper>
    </div>
  </article>


</template>

<script lang="ts">
import { Swiper, SwiperSlide } from "swiper/vue";

import "swiper/swiper-bundle.min.css";
import "swiper/swiper.min.css";
import "swiper/components/navigation/navigation.min.css";
import SwiperCore, { Navigation } from "swiper";
import { defineComponent, PropType } from "vue";

import HeroArticle from "@/types/HeroArticle";

SwiperCore.use([Navigation]);
export default defineComponent({
  props: {
    articles: {
      required: true,
      type: Array as PropType<HeroArticle[]>
    }
  },
  components: {
    Swiper,
    SwiperSlide
  }

});
</script>

<style scoped lang="sass">
@import "../assets/sass/helps/mixins"


.hero-slider
  height: 100%
  width: 100%

.hero-swiper
  width: 100%
  height: 100%

.hero-swiper
  &.swiper-button-next
    top: 10% !important

.article
  background-color: #000
  max-width: 50%
  height: 50vh

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
    @extend %subtitle-one


.button
  padding: 16px 34px
  font-weight: bold
  font-size: 1rem
  position: absolute
  bottom: 62px
  left: 50px
  margin-top: 34px

  &__prev,
  &__next
    width: 64px
    height: 64px

  &__prev
    background-color: #fff

  &__next
    background-color: #000

</style>