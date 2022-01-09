<template>
  <div class="relative">
    <button class="z-10 relative flex items-center focus:outline-none select-none" @click="open = !open">
      <slot name="button"></slot>
    </button>

    <!-- to close when clicked on space around it-->
    <button class="fixed inset-0 h-full w-full cursor-default focus:outline-none" v-if="open" @click="open = false" tabindex="-1"></button>

    <!--dropdown menu-->
    <transition name="fade">
      <div class="absolute shadow-lg border w-48 rounded py-1 px-2 text-sm mt-4 bg-white z-10" :class="placement === 'right' ? 'right-0' : 'left-0'" v-if="open">
        <slot name="content"></slot>
      </div>
    </transition>
  </div>
</template>

<script>
  export default {
    name: "ar-dropdown",
    data() {
      return {
        open: false,
      };
    },
    props: {
      placement: {
        type: String,
        default: "right",
        validator: (value) => ["right", "left"].indexOf(value) !== -1,
      },
    },
    mounted() {
      const onEscape = (e) => {
        if (e.key === "Esc" || e.key === "Escape") {
          this.open = false;
        }
      };

      document.addEventListener("keydown", onEscape);

      this.$once("hook:beforeDestroy", () => {
        document.removeEventListener("keydown", onEscape);
      });
    },
  };
</script>

<style lang="scss" scoped>
.fade-enter-active, .fade-leave-active {
  transition: all .2s ease;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
  transform: scale(0.95);
}
</style>