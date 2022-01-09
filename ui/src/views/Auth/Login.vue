<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full">
      <div>
        <!-- <img class="mx-auto h-12 w-auto" src="" alt="DevSpace.app"> -->
        <h2 class="mt-6 text-center text-3xl leading-9 font-extrabold text-gray-900">
          Sign in to your account
        </h2>
        <p class="mt-2 text-center text-sm leading-5 text-gray-600">
          Or
          <router-link :to="{ name: 'Auth/Register' }" class="font-medium text-indigo-600 hover:text-indigo-500 focus:outline-none focus:underline transition ease-in-out duration-150">
            create a new account
          </router-link>
        </p>
      </div>
      <div class="rounded-md bg-white shadow-lg p-8 mt-4">
        <form caction="#" method="POST">
          <input type="hidden" name="remember" value="true">

          <label for="email" class="block text-base leading-8 font-medium text-gray-700">Username</label>
          <input aria-label="Email address" name="email" type="email" required class="appearance-none shadow block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5" placeholder="or email address">
          
          <label for="password" class="block text-base leading-8 font-medium text-gray-700 mt-4">Password</label>
          <input aria-label="Password" name="password" type="password" required class="appearance-none shadow block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5">

          <div class="mt-6 flex items-center justify-between">
            <div class="flex items-center">
              <input id="remember_me" type="checkbox" class="form-checkbox shadow h-4 w-4 text-indigo-600 transition duration-150 ease-in-out">
              <label for="remember_me" class="ml-2 block text-sm leading-5 text-gray-900">
                Remember me
              </label>
            </div>

            <div class="text-sm leading-5">
              <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500 focus:outline-none focus:underline transition ease-in-out duration-150">
                Forgot your password?
              </a>
            </div>
          </div>

          <div class="mt-6">
            <button type="submit" class="group relative shadow-lg w-full flex justify-center py-2 px-4 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-300 ease-in-out">
              Sign in
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
  name: 'Auth/Login',
  components: {
  }
}
</script>
