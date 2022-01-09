<template>
  <common-layout tab="sessions">
    <header>
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold leading-tight text-gray-900">
          Sessions
        </h1>
      </div>
    </header>

    <main>
      <div class="max-w-7xl mx-auto sm:px-6 lg:px-8">
        <div class="rounded-md bg-white shadow-lg p-6">
          You have no active sessions at this current moment
        </div>
      </div>
    </main>
  </common-layout>
</template>

<style lang="scss" scoped>

</style>

<script>
// @ is an alias to /src
import CommonLayout from '@/components/CommonLayout.vue'

export default {
  name: 'Sessions',
  components: {
    CommonLayout
  },
  data () {
    return {
      publicPath: process.env.BASE_URL
    }
  }
}
</script>
