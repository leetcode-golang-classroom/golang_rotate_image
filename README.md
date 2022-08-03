# golang_rotate_image

You are given an `n x n` 2D `matrix` representing an image, rotate the image by **90** degrees (clockwise).

You have to rotate the image **[in-place](https://en.wikipedia.org/wiki/In-place_algorithm)**, which means you have to modify the input 2D matrix directly. **DO NOT** allocate another 2D matrix and do the rotation.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/08/28/mat1.jpg](https://assets.leetcode.com/uploads/2020/08/28/mat1.jpg)

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/08/28/mat2.jpg](https://assets.leetcode.com/uploads/2020/08/28/mat2.jpg)

```
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

```

**Constraints:**

- `n == matrix.length == matrix[i].length`
- `1 <= n <= 20`
- `1000 <= matrix[i][j] <= 1000`

## 解析

給定一個矩陣 matrix 

要求在不使用其他額外存儲空間下

把 matrix 順時針方向 旋轉 90 度

要求寫出這樣的演算法

因為要求不能使用額外空間

先想想哪些操作可以不使用額外空間

代表這些操作一次需要對換成對的位置

作法1：  順時針轉 90 度 = 先把矩陣做水平翻轉  然後再針對右上到左下的對角線反轉

![](https://i.imgur.com/OaaN74Z.png)
![](https://i.imgur.com/AafcZBK.png)
![](https://i.imgur.com/HXLQOXj.png)

這樣的時間複雜度是 O($n^2$)

空間複雜度是 O(1)

作法2 ：

仔細去看 選順時針轉 90 度， 剛好是把 矩陣 4 個部份的數值做位置互換如下

![](https://i.imgur.com/n7AoCAX.png)
![](https://i.imgur.com/0xFnQDt.png)

四個象限的座標對應如下:

假設座標再第1象限是 (i,j)

其第2象限是 (j, n-1-i) 等於方向是 x, y 互換, 然後 把y 方向做補數

其第3象限是 (n-1-i, n-1-j) 等於方向是 把x, y 方向做補數

其第4象限是 (n-1-j, i) 等於方向是 x, y 互換, 然後 把x 方向做補數

替換的演算法如下

i = 0..Math.ceil(n/2), j= 0..Math.floor(n/2)

temp = matrix[i][j]
matrix[i][j] = matrix[n-1-j][i]
matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
matrix[n-1-i][n-1-j]= matrix[j][n-1-i]
matrix[j][n-1-i] = temp

這樣的時間複雜度也是 O($n^2$)

然而運算的時間比前一個少了一半因為只需要處理 1/4 的面積

前一個需要對換兩次1/2 的面積

空間複雜度也是 O(1)

## 程式碼
```go
package sol

func rotate(matrix [][]int) {
	n := len(matrix)
	small := n / 2
	large := small
	if n%2 != 0 {
		large = small + 1
	}
	for i := 0; i < large; i++ {
		for j := 0; j < small; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}
```
## 困難點

1. 需要找出對換值的位置關係

## Solve Point

- [x]  算出 if n % 2 == 0 , max = n/2, min = n/2 否則 max = n/2 + 1, min = n/2
- [x]  對 i = 0..max, j= 0..min 做以下替換
- [x]  temp = matrix[i][j], matrix[i][j] = matrix[n-1-j][i], matrix[n-1-j][i] = matrix[n-1-i][n-1-j], matrix[n-1-i][n-1-j]= matrix[j][n-1-i], matrix[j][n-1-i] = temp