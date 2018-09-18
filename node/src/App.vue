<template>
    <div class="container" id="app">
        <div class="app">
            <h1>Static Site</h1>
            <at-button @click="handleClick">Click</at-button>
        </div>
    </div>
</template>

<script>
    export default {
      methods: {
        handleClick() {
          this.$Message('Tester')
        }
      }
    }
</script>

<style{{#sass}} lang="scss"{{/sass}}>
    .container {
        width: 100%;
        height: 100%;
    }
    .app {
        text-align: center;
    }
    h1 {
        font-weight: normal;
    }
</style>