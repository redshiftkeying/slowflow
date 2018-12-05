<template>
    <div class="accept-container">
      <div id="placeholder"></div>
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

  },
  methods:{
  }
}
new DocsAPI.DocEditor("placeholder", {
    "document": {
        "fileType": "xlsx",
        "key": "Khirz6zTPdfd7",
        "title": "Example Document Title.docx",
        "url": "http://127.0.0.1:8090/static/PD-004.xlsx"
    },
    "documentType": "text"
});
</script>
<style scoped>


</style>

