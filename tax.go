/**
*copyright on elvin @2019
*used to calucate tax and clean income.if you have any useful idea. you are welcome
*to contribute this project.
*/
package main

import (
   "fmt"
   "os"
   "flag"
)
func main() {
	  argcList := os.Args
	  if len(argcList) < 2 {
	     fmt.Printf("Args: .exe -beforeTax(税前工资) -unTax(免税额度，可省略)\n")
		 return
	  }
	  
	  var totalMoney float64 = 0.
      var unTax float64 = 0.
	  flag.Float64Var(&totalMoney,"beforeTax",0.,"税前总工资")
	  flag.Float64Var(&unTax,"unTax",5000.0,"基本征收起点5000+专项附加扣除，比如租房子，养孩子，老人")
	  flag.Parse()
	  
      var get1 float64 = 0
      var get2 float64 = 210
      var get3 float64 = 1410
      var get4 float64 = 2660
      var get5 float64 = 4410
      var get6 float64 = 7160
      var get7 float64 = 15160
      var money float64 = 0
      if totalMoney <= 5000 {
           money = totalMoney
      }else if totalMoney>5000 && totalMoney <=8000{
           money = totalMoney -(totalMoney - unTax)*0.03 + get1
      }else if totalMoney>8000 && totalMoney <=17000{
           money = totalMoney -(totalMoney - unTax)*0.1 + get2
      }else if totalMoney>17000 && totalMoney <=30000{
           money = totalMoney -(totalMoney - unTax)*0.2 + get3
      }else if totalMoney>30000 && totalMoney <=40000{
           money = totalMoney -(totalMoney - unTax)*0.25 + get4
      }else if totalMoney>40000 && totalMoney <=60000{
           money = totalMoney -(totalMoney - unTax)*0.3 + get5
      }else if totalMoney>60000 && totalMoney <=85000{
           money = totalMoney -(totalMoney - unTax)*0.35 + get6
      }else{
           money = totalMoney -(totalMoney - unTax)*0.45 + get7
      }
      fmt.Printf("你所承担的个人所得税是：%f!!!!!\n",totalMoney - money)
      fmt.Printf("税后净收入是: %f!!!!!\n",money)
	  fmt.Printf("这个净收入指的是没有扣除，社保和公积金的收入，当然你如果将社保和公积金看作是收入的话......")
}