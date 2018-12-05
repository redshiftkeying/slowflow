<template>
    <div class="accept-container">
      <iframe id="show-iframe"  frameborder=0 name="showHere" scrolling=auto :src="link"></iframe>
    </div>
</template>

<script>
export default {
  data () {
    return {
      link: 'http://192.168.31.211:8080/fine/ReportServer?reportlet=PDF-03.cpt&op=write&id=6d5d2078e4234fec94cf7efae6f3d012'
    }
  },
  mounted(){
    const oIframe = document.getElementById('show-iframe');
    const deviceWidth = document.documentElement.clientWidth;
    const deviceHeight = document.documentElement.clientHeight;
    oIframe.style.width = deviceWidth + 'px';
    oIframe.style.height = deviceHeight + 'px';
  },
  methods:{
  }
}

</script>
<style scoped>


</style>

