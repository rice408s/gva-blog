<template>
  <div class="main">
    <Layout>
      <Header>
        <NavBar />
      </Header>
      <Layout>

        <Content>
          <RouterView/>
        </Content>

      </Layout>
      <Footer><FootBar/></Footer>
    </Layout>
  </div>
</template>

<script setup>
import  NavBar  from './components/NavBar.vue';
import FootBar from './components/FootBar.vue'
</script>


<style scoped>

.RouterView {
  background-color: #f0f2f5;
  height: 200%;
  width: 200%;
  display: flex;
  margin: auto;
  /* justify-content: center; */
  align-items: center;
  display: block;
}

.Content {
  /* background-color: #f0f2f5; */
  height: 200%;
  width: 200%;
  display: block;
  justify-content: center;
  align-items: center;
}

.layout {
  /* 最小高度占据整个视窗 */
  min-height: 100vh;
  max-width: 200%;
  width: 200%;
  display: flex;
  flex-direction: column;
}

.FootBar {
  /* background-color: #f0f2f5; */
  height: 69px;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
}

</style>
