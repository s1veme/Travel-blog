<template>
  <div class="who">
    <div class="about">
      <div class="about__img">
        <img src="../assets/profile.jpg" alt="Jaspreet Bhamrai">
      </div>
      <h2 class="about__title heading-three">Jaspreet Bhamrai</h2>
      <p class="about__text paragraph">
        For as long as I can remember I’ve been obsessed with the idea of travel. I was always that person who was
        forever daydreaming of foreign lands and unfamiliar cultures; coming up with travel itineraries that would
        challenge my perceptions and help me gain a deeper understanding of the world.
      </p>
      <p class="about__follow subtitle-two">
        Follow me
      </p>
      <div class="about__socials">
        <div class="about__socials-item" v-for="(item, index) in 4" :key="index">

        </div>
      </div>
    </div>
    <div class="cities">
      <div class="cities-item">
        <p>
          Tokyo
        </p>
      </div>
      <div class="cities-item">
        <p>
          Rome
        </p>
      </div>
      <div class="cities-item">
        <p>
          San Francisco
        </p>
      </div>
      <div class="cities-item">
        <p>
          San Jose
        </p>
      </div>
    </div>
    <div class="subscribe">
      <p class="subscribe__title heading-three">Newsletter</p>
      <p class="subscribe__subtitle paragraph">Subscribe to receive exclusive content updates, travel & photo tips!</p>
      <form>
        <label class="subscribe__label who__label" for="subscribeEmail">Email address</label>
        <input type="text" id="subscribeEmail" placeholder="example@|" class="subscribe__input who__input paragraph">
        <button type="button" class="button__black subscribe__button">Subscribe</button>
      </form>
    </div>
    <div class="where">
      <form>
        <p class="subscribe__title heading-three">Where to next?</p>
        <p class="subscribe__subtitle paragraph">
          SubscTher’s a wide world waiting for you
        </p>
        <label for="destination">Destination name</label>
        <input class="where__destination who__input paragraph" id="destination" type="text" placeholder="Japan">
        <input class="who__input paragraph" type="text" placeholder="Check-in date">
        <input class="who__input paragraph" type="text" placeholder="Check-out date">
        <button type="button" class="button__black subscribe__button">Read more</button>
      </form>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style scoped lang="sass">
@import "../assets/sass/helps/mixins"
.cities
  margin-top: 82px

  &-item
    @extend %heading-three
    background-color: #fff
    width: 100%
    height: 125px
    margin-bottom: 0.5rem
    color: #000
    display: flex
    flex-direction: column
    align-items: center
    justify-content: center

    p
      position: relative
      margin-top: 10px

      &::before
        content: ''
        position: absolute
        top: -10px
        left: 50%
        width: 64px
        height: 4px
        background-color: #000
        transform: translateX(-50%)

.who
  border: 2px solid #fff
  max-height: 482px

.about
  text-align: center
  padding: 30px
  position: relative
  height: 100%

  &__socials
    display: flex
    align-items: center
    justify-content: space-between
    gap: 10px
    position: absolute
    bottom: -31px
    left: 50%
    transform: translateX(-50%)

    &-item
      width: 62px
      height: 62px
      background-color: #fff

  &__follow
    font-weight: bold
    margin-bottom: 2rem
    margin-top: 4.8rem

  &__title
    margin-top: 65px

  &__text
    color: #949799
    font-size: 16px

  &__img
    height: 152px
    width: 152px
    margin: 0 auto
    position: absolute
    top: 0
    left: 50%
    transform: translate(-50%, -50%)

    img
      height: 100%
      width: 100%
      object-fit: cover
      border-radius: 50%


.who
  width: 40%

  &__label
    font-size: 14px
    margin-bottom: 7px

  &__input
    background-color: transparent
    outline: none
    border: none
    border-bottom: 2px solid #ffffff
    width: 100%
    color: #fff
    margin-bottom: 2rem

    &::placeholder
      color: #fff

.subscribe
  border: 2px solid #fff
  padding: 30px

  &__button
    width: 100%
    height: 55px
    font-weight: bold

.where
  border: 2px solid #fff
  padding: 30px
</style>