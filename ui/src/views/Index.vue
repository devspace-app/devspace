<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full">
      <div>
        <h2 class="mt-6 text-center text-3xl leading-9 font-extrabold text-gray-900">
          Welcome to DevSpace.app
        </h2>
        <p class="mt-2 text-center text-sm leading-5 text-gray-600">
          temporary index page
        </p>
      </div>
      <div class="rounded-md bg-white shadow-lg p-8 mt-4">
        <router-link :to="{ name: 'Auth/Login' }" class="group relative shadow-lg w-full flex justify-center py-2 px-4 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-300 ease-in-out">
          Sign in
        </router-link>
        <router-link :to="{ name: 'Dashboard' }" class="mt-4 group relative shadow-lg w-full flex justify-center py-2 px-4 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-300 ease-in-out">
          Dashboard
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>


export default {
  name: 'Index',
  components: {
  }
}
</script>
