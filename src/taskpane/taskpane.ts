/*
 * Copyright (c) Microsoft Corporation. All rights reserved. Licensed under the MIT license.
 * See LICENSE in the project root for license information.
 */

/* global console, document, Excel, Office */
// const pg = require('pg');
// const cs = 'postgres://postgres:wolfe;@localhost:5432/jauntyj';

// const client = new pg.Client(cs);
// client.connect();
Office.onReady((info) => {
  if (info.host === Office.HostType.Excel) {
    document.getElementById("sideload-msg").style.display = "none";
    document.getElementById("app-body").style.display = "flex";
    document.getElementById("run").onclick = run;
  }
});

export async function run() {

  
  try {
    await Excel.run(async (context) => {
      /**
       * Insert your Excel code here
       */
      const range: Excel.Range = context.workbook.getSelectedRange();
      
      

      // Read the range address
      range.load("address");
      //range.values = [["test1"]];

      // Update the fill color
      range.format.fill.color = "yellow";
      console.log("testtest")
      await webRequest(context);

      

      



      await context.sync();
      console.log(`The range address was ${range.address}.`);
    });
  } catch (error) {
    console.error(error);
  }

  
}
var webRequest = async (context) => {
  



  try {

    // Make a new HTTP Request.
    var xhr = new XMLHttpRequest();

    // Define the Method and URL.
    xhr.open('GET', `https://localhost:8080/hello`);



    // Print the Text to the Console.
    xhr.onload = function (e) {
      console.log(this.responseText);
      console.log(e);
      context.workbook.getSelected
      var range: Excel.Range = context.workbook.getSelectedRange();
      range.values=[[this.responseText]]
      return context.sync()
    }

    // Send the Request.
    xhr.send();

  } catch (error) {
    console.error(error);
  }

}

// async function webRequest(range: Excel.Range) {
//   //let url = "https://localhost:8080/hello";
//   let url = "https://httpbin.org/ip"
//   var xhttp = new XMLHttpRequest();
//   xhttp.addEventListener("progress", updateProgress)
//   xhttp.open("GET", url);
//   try{
//   range.values = [['testest']]
//   range.values = [[xhttp.readyState]]
//   xhttp.onreadystatechange= function() {
//     if (true) {
//       range.values =
//       [[this.readyState]];
//     }

//   }
//   xhttp.send(null);
//   range.values = [[xhttp.readyState]]

// } catch(error){
//   range.values = [[error]]
// }
// function updateProgress(xhttp){
//   [[xhttp.readyState]];

// }
// }
// async function getStarCount() {

//   const url = "https://localhost:8080/hello";

//   let xhttp = new XMLHttpRequest();

//   return new Promise(function(resolve, reject) {
//     xhttp.onreadystatechange = function() {
//       if (xhttp.readyState !== 4) return;

//       if (xhttp.status == 200) {
//         resolve(xhttp.responseText);
//       } else {
//         resolve({
//           status: xhttp.status,

//           statusText: xhttp.statusText
//         });
//       }
//     };

//     xhttp.open("GET", url, true);

//     xhttp.send();
//   });
// }


