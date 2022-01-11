# **Playground**
[MatrixOne Playground](https://playground.matrixorigin.io/?tutorial=SSB-test-with-matrixone&step=1) allow you to try SQL statements and explore features of MatrixOne instantly from your web browser with interactive tutorials.  

* For more guided information about Playground,you can read this page detailly.  
* If you have some questions about SQL,you can also see [SQL Reference](../Reference/SQL-Reference/Data-Definition-Statements/create-database.md).  
* For text and code tutorials about **SSB** test,you can see [Tutorial](../Get-Started/Tutorial/SSB-test-with-matrixone.md).

## **Limitations**

You can only operate in **read-only** mode in MatixOrigin Playground，so DDL commands and part of DML commands which may change the data are not available.The detail limitations are shown below:

* **DDL** commands are not available:  
```create/drop table``` , ```truncate``` , ```update``` , ```set``` ,```use```  
  
* Part of **DML** commands are not available：  
```insert``` , ```replace``` , ```delete```,```select into ```  

* ```commit``` is not available

* ```call``` is not available

* **max_result_rows**=2000  
## **Examples**
Now,You can follow the tutorials to start Playground in the example below.  
[**Click here to turn to Playground**](https://playground.matrixorigin.io/?tutorial=SSB-test-with-matrixone&step=1)  
### ** Test Preperations**  

This tutorial walks you through the most popular **Star Schema Benchmark（SSB）**Test SQL statements with MatrixOne. To better experience MatrixOne features and performance, test queries in this tutorial will run without filters.  
Before you started, the test datasets have been pre-loaded in  database `ssb`. To list available tables in the database you can query :

```
SHOW TABLES；
```

!!! note  "Tips"
    You can click on the command to copy and **run** it in the terminal on the right.  
    Then you can run the queries on SSB datasets, click **Continue**.

### **Run Query Command**
Now, You can query the table with SQL commands we provide:  

#### **1. Run Q1.1 query**

```
select sum(lo_revenue) as revenue
from lineorder join dates on lo_orderdate = d_datekey；
```

#### **2. Run Q1.2 query**

```
select sum(lo_revenue) as revenue
from lineorder
join dates on lo_orderdate = d_datekey；
```

#### **3. Run Q1.3 query**

```
select sum(lo_revenue) as revenue
from lineorder
join dates on lo_orderdate = d_datekey
```

#### **4. Run Q2.1 query**

```
select sum(lo_revenue) as lo_revenue, d_year, p_brand
from lineorder
join dates on lo_orderdate = d_datekey
join part on lo_partkey = p_partkey
join supplier on lo_suppkey = s_suppkey
group by d_year, p_brand
order by d_year, p_brand;
```


#### **5. Run Q2.2 query**

```
select sum(lo_revenue) as lo_revenue, d_year, p_brand
from lineorder
join dates on lo_orderdate = d_datekey
join part on lo_partkey = p_partkey
join supplier on lo_suppkey = s_suppkey
group by d_year, p_brand
order by d_year, p_brand;
```


#### **6. Run Q2.3 query**

```
select sum(lo_revenue) as lo_revenue, d_year, p_brand
from lineorder
join dates on lo_orderdate = d_datekey
join part on lo_partkey = p_partkey
join supplier on lo_suppkey = s_suppkey
group by d_year, p_brand
order by d_year, p_brand;
```


#### **7. Run Q3.1 query**

```
select c_nation, s_nation, d_year, sum(lo_revenue) as lo_revenue
from lineorder
join dates on lo_orderdate = d_datekey
join customer on lo_custkey = c_custkey
join supplier on lo_suppkey = s_suppkey
group by c_nation, s_nation, d_year
order by d_year asc, lo_revenue desc;
```

#### **8. Run Q3.2 query**

```
select c_city, s_city, d_year, sum(lo_revenue) as lo_revenue
from lineorder
join dates on lo_orderdate = d_datekey
join customer on lo_custkey = c_custkey
join supplier on lo_suppkey = s_suppkey
group by c_city, s_city, d_year
order by d_year asc, lo_revenue desc;
```


#### **9. Run Q3.3 query**

```
select c_city, s_city, d_year, sum(lo_revenue) as lo_revenue
from lineorder
join dates on lo_orderdate = d_datekey
join customer on lo_custkey = c_custkey
join supplier on lo_suppkey = s_suppkey
group by c_city, s_city, d_year
order by d_year asc, lo_revenue desc;
```

#### **10. Run Q3.4 query**

```
select c_city, s_city, d_year, sum(lo_revenue) as lo_revenue
from lineorder
join dates on lo_orderdate = d_datekey
join customer on lo_custkey = c_custkey
join supplier on lo_suppkey = s_suppkey
group by c_city, s_city, d_year
order by d_year asc, lo_revenue desc;
```

#### **11. Run Q4.1 query**

```
select d_year, c_nation, sum(lo_revenue) - sum(lo_supplycost) as profit
from lineorder
join dates on lo_orderdate = d_datekey
join customer on lo_custkey = c_custkey
join supplier on lo_suppkey = s_suppkey
join part on lo_partkey = p_partkey
group by d_year, c_nation
order by d_year, c_nation;
```

#### **12. Run Q4.2 query**

```
select d_year, s_nation, p_category, sum(lo_revenue) - sum(lo_supplycost) as profit
from lineorder
join dates on lo_orderdate = d_datekey
join customer on lo_custkey = c_custkey
join supplier on lo_suppkey = s_suppkey
join part on lo_partkey = p_partkey
group by d_year, s_nation, p_category
order by d_year, s_nation, p_category;
```

#### **13. Run Q4.3 query**

```
select d_year, s_city, p_brand, sum(lo_revenue) - sum(lo_supplycost) as profit
from lineorder
join dates on lo_orderdate = d_datekey
join customer on lo_custkey = c_custkey
join supplier on lo_suppkey = s_suppkey
join part on lo_partkey = p_partkey
group by d_year, s_city, p_brand
order by d_year, s_city, p_brand;
```

## **Learn More**
This page describes the features, limitations, and examples of Playground. For information on other options that are available when trying out MatrixOne, see the following:

* [Install MatrixOne](install-standalone-matrixone.md)
* [What‘s New](../Overview/what's-new.md)