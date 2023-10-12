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
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}


.FootBar {
  background-color: #f0f2f5;
  height: 69px;
  display: flex;
  justify-content: center;
  align-items: center;
}

</style>
