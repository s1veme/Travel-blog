<template>
  <div class="categories-slider">
    <div class="categories-slider__container">
      <h2 class="heading-three title">Explore by category</h2>
      <div class="categories-slider-block">
        <swiper
          :slidesPerView="8"
          :spaceBetween="30"

          :navigation="true"
          class="mySwiper w-1/4"
        >
          <swiper-slide
            class="categories-slider-slide"
            v-for="(item, index) in 10"
            :key="index"
            >{{ index }}
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </div>
</template>

<script>
import { Swiper, SwiperSlide } from "swiper/vue";

import "swiper/swiper-bundle.min.css";
import "swiper/swiper.min.css";
import "swiper/components/navigation/navigation.min.css";
import SwiperCore, { Navigation } from "swiper";

SwiperCore.use([Navigation]);
export default {
  components: {
    Swiper,
    SwiperSlide
  }
};
</script>
<style scoped lang="sass">
@import "../assets/sass/helps/mixins"

.categories-slider-block
  margin-bottom: 120px
  margin-right: 0
  width: auto
  height: 315px

.categories-slider
  &-slide
    height: 250px
    width: 250px
    background-color: #fff
    color: #000
    text-align: center
    display: flex
    flex-direction: column
    align-items: center
    justify-content: center
    transition: all 300ms ease-in-out

    &:hover
      background-color: #7d7d7d
      transform: scale(1.2)
      color: #fff
      height: 315px

.title
  margin-bottom: 60px
  margin-top: 70px

</style>
