<script lang="ts">
  import "../app.css";
  let { children } = $props();
</script>

<main class="w-full h-screen lg:px-[25vw] p-6">
  <div class="h-full">
    {@render children()}
  </div>
</main>
