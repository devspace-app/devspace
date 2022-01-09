<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full">
      <div>
        <!-- <img class="mx-auto h-12 w-auto" src="" alt="DevSpace.app"> -->
        <h2 class="mt-6 text-center text-3xl leading-9 font-extrabold text-gray-900">
          Create a new account
        </h2>
        <p class="mt-2 text-center text-sm leading-5 text-gray-600">
          Or
          <router-link :to="{ name: 'Auth/Login' }" class="font-medium text-indigo-600 hover:text-indigo-500 focus:outline-none focus:underline transition ease-in-out duration-150">
            login to your existing account
          </router-link>
        </p>
      </div>
      <div class="rounded-md bg-white shadow-lg p-8 mt-4">
        <form caction="#" method="POST">
          <input type="hidden" name="remember" value="true">

          <label for="username" class="block text-base leading-8 font-medium text-gray-700">Username</label>
          <input aria-label="Username" name="username" type="email" required class="appearance-none shadow block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5">
          
          <label for="email" class="block text-base leading-8 font-medium text-gray-700 mt-4">Email</label>
          <input aria-label="Email address" name="email" type="email" required class="appearance-none shadow block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5">

          <label for="password" class="block text-base leading-8 font-medium text-gray-700 mt-4">Password</label>
          <input aria-label="Password" name="password" type="password" required class="appearance-none shadow block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5">

          <label for="password_confirm" class="block text-base leading-8 font-medium text-gray-700 mt-4">Confirm Password</label>
          <input aria-label="Confirm Password" name="password_confirm" type="password" required class="appearance-none shadow block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5">

          <div class="mt-6">
            <button type="submit" class="group relative shadow-lg w-full flex justify-center py-2 px-4 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-300 ease-in-out">
              Register
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'Auth/Register',
  components: {
  }
}
</script>
