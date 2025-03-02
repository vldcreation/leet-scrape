package testdata

import "github.com/vldcreation/leet-scrape/v2/domain/entity"

var QuestionWithImg = entity.Question{
	TitleSlug: "convert-sorted-array-to-binary-search-tree",
	Content:   "<p>Given an integer array <code>nums</code> where the elements are sorted in <strong>ascending order</strong>, convert <em>it to a <strong>height-balanced</strong> binary search tree</em>.</p>\n\n<p>A <strong>height-balanced</strong> binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.</p>\n\n<p>&nbsp;</p>\n<p><strong>Example 1:</strong></p>\n<img alt=\"\" src=\"https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg\" style=\"width: 302px; height: 222px;\" />\n<pre>\n<strong>Input:</strong> nums = [-10,-3,0,5,9]\n<strong>Output:</strong> [0,-3,9,-10,null,5]\n<strong>Explanation:</strong> [0,-10,5,null,-3,null,9] is also accepted:\n<img alt=\"\" src=\"https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg\" style=\"width: 302px; height: 222px;\" />\n</pre>\n\n<p><strong>Example 2:</strong></p>\n<img alt=\"\" src=\"https://assets.leetcode.com/uploads/2021/02/18/btree.jpg\" style=\"width: 342px; height: 142px;\" />\n<pre>\n<strong>Input:</strong> nums = [1,3]\n<strong>Output:</strong> [3,1]\n<strong>Explanation:</strong> [1,3] and [3,1] are both a height-balanced BSTs.\n</pre>\n\n<p>&nbsp;</p>\n<p><strong>Constraints:</strong></p>\n\n<ul>\n\t<li><code>1 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>\n\t<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>\n\t<li><code>nums</code> is sorted in a <strong>strictly increasing</strong> order.</li>\n</ul>\n",
	CodeSnippets: []entity.CodeSnippet{
		{
			Lang:     "C++",
			LangSlug: "cpp",
			Code:     "/**\n * Definition for a binary tree node.\n * struct TreeNode {\n *     int val;\n *     TreeNode *left;\n *     TreeNode *right;\n *     TreeNode() : val(0), left(nullptr), right(nullptr) {}\n *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}\n *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}\n * };\n */\nclass Solution {\npublic:\n    TreeNode* sortedArrayToBST(vector<int>& nums) {\n        \n    }\n};",
		},
		{
			Lang:     "Java",
			LangSlug: "java",
			Code:     "/**\n * Definition for a binary tree node.\n * public class TreeNode {\n *     int val;\n *     TreeNode left;\n *     TreeNode right;\n *     TreeNode() {}\n *     TreeNode(int val) { this.val = val; }\n *     TreeNode(int val, TreeNode left, TreeNode right) {\n *         this.val = val;\n *         this.left = left;\n *         this.right = right;\n *     }\n * }\n */\nclass Solution {\n    public TreeNode sortedArrayToBST(int[] nums) {\n        \n    }\n}",
		},
		{
			Lang:     "Python",
			LangSlug: "python",
			Code:     "# Definition for a binary tree node.\n# class TreeNode(object):\n#     def __init__(self, val=0, left=None, right=None):\n#         self.val = val\n#         self.left = left\n#         self.right = right\nclass Solution(object):\n    def sortedArrayToBST(self, nums):\n        \"\"\"\n        :type nums: List[int]\n        :rtype: TreeNode\n        \"\"\"\n        ",
		},
		{
			Lang:     "Python3",
			LangSlug: "python3",
			Code:     "# Definition for a binary tree node.\n# class TreeNode:\n#     def __init__(self, val=0, left=None, right=None):\n#         self.val = val\n#         self.left = left\n#         self.right = right\nclass Solution:\n    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:\n        ",
		},
		{
			Lang:     "C",
			LangSlug: "c",
			Code:     "/**\n * Definition for a binary tree node.\n * struct TreeNode {\n *     int val;\n *     struct TreeNode *left;\n *     struct TreeNode *right;\n * };\n */\n\n\nstruct TreeNode* sortedArrayToBST(int* nums, int numsSize){\n\n}",
		},
		{
			Lang:     "C#",
			LangSlug: "csharp",
			Code:     "/**\n * Definition for a binary tree node.\n * public class TreeNode {\n *     public int val;\n *     public TreeNode left;\n *     public TreeNode right;\n *     public TreeNode(int val=0, TreeNode left=null, TreeNode right=null) {\n *         this.val = val;\n *         this.left = left;\n *         this.right = right;\n *     }\n * }\n */\npublic class Solution {\n    public TreeNode SortedArrayToBST(int[] nums) {\n        \n    }\n}",
		},
		{
			Lang:     "JavaScript",
			LangSlug: "javascript",
			Code:     "/**\n * Definition for a binary tree node.\n * function TreeNode(val, left, right) {\n *     this.val = (val===undefined ? 0 : val)\n *     this.left = (left===undefined ? null : left)\n *     this.right = (right===undefined ? null : right)\n * }\n */\n/**\n * @param {number[]} nums\n * @return {TreeNode}\n */\nvar sortedArrayToBST = function(nums) {\n    \n};",
		},
		{
			Lang:     "Ruby",
			LangSlug: "ruby",
			Code:     "# Definition for a binary tree node.\n# class TreeNode\n#     attr_accessor :val, :left, :right\n#     def initialize(val = 0, left = nil, right = nil)\n#         @val = val\n#         @left = left\n#         @right = right\n#     end\n# end\n# @param {Integer[]} nums\n# @return {TreeNode}\ndef sorted_array_to_bst(nums)\n    \nend",
		},
		{
			Lang:     "Swift",
			LangSlug: "swift",
			Code:     "/**\n * Definition for a binary tree node.\n * public class TreeNode {\n *     public var val: Int\n *     public var left: TreeNode?\n *     public var right: TreeNode?\n *     public init() { self.val = 0; self.left = nil; self.right = nil; }\n *     public init(_ val: Int) { self.val = val; self.left = nil; self.right = nil; }\n *     public init(_ val: Int, _ left: TreeNode?, _ right: TreeNode?) {\n *         self.val = val\n *         self.left = left\n *         self.right = right\n *     }\n * }\n */\nclass Solution {\n    func sortedArrayToBST(_ nums: [Int]) -> TreeNode? {\n        \n    }\n}",
		},
		{
			Lang:     "Go",
			LangSlug: "golang",
			Code:     "/**\n * Definition for a binary tree node.\n * type TreeNode struct {\n *     Val int\n *     Left *TreeNode\n *     Right *TreeNode\n * }\n */\nfunc sortedArrayToBST(nums []int) *TreeNode {\n    \n}",
		},
		{
			Lang:     "Scala",
			LangSlug: "scala",
			Code:     "/**\n * Definition for a binary tree node.\n * class TreeNode(_value: Int = 0, _left: TreeNode = null, _right: TreeNode = null) {\n *   var value: Int = _value\n *   var left: TreeNode = _left\n *   var right: TreeNode = _right\n * }\n */\nobject Solution {\n    def sortedArrayToBST(nums: Array[Int]): TreeNode = {\n        \n    }\n}",
		},
		{
			Lang:     "Kotlin",
			LangSlug: "kotlin",
			Code:     "/**\n * Example:\n * var ti = TreeNode(5)\n * var v = ti.`val`\n * Definition for a binary tree node.\n * class TreeNode(var `val`: Int) {\n *     var left: TreeNode? = null\n *     var right: TreeNode? = null\n * }\n */\nclass Solution {\n    fun sortedArrayToBST(nums: IntArray): TreeNode? {\n        \n    }\n}",
		},
		{
			Lang:     "Rust",
			LangSlug: "rust",
			Code:     "// Definition for a binary tree node.\n// #[derive(Debug, PartialEq, Eq)]\n// pub struct TreeNode {\n//   pub val: i32,\n//   pub left: Option<Rc<RefCell<TreeNode>>>,\n//   pub right: Option<Rc<RefCell<TreeNode>>>,\n// }\n// \n// impl TreeNode {\n//   #[inline]\n//   pub fn new(val: i32) -> Self {\n//     TreeNode {\n//       val,\n//       left: None,\n//       right: None\n//     }\n//   }\n// }\nuse std::rc::Rc;\nuse std::cell::RefCell;\nimpl Solution {\n    pub fn sorted_array_to_bst(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {\n        \n    }\n}",
		},
		{
			Lang:     "PHP",
			LangSlug: "php",
			Code:     "/**\n * Definition for a binary tree node.\n * class TreeNode {\n *     public $val = null;\n *     public $left = null;\n *     public $right = null;\n *     function __construct($val = 0, $left = null, $right = null) {\n *         $this->val = $val;\n *         $this->left = $left;\n *         $this->right = $right;\n *     }\n * }\n */\nclass Solution {\n\n    /**\n     * @param Integer[] $nums\n     * @return TreeNode\n     */\n    function sortedArrayToBST($nums) {\n        \n    }\n}",
		},
		{
			Lang:     "TypeScript",
			LangSlug: "typescript",
			Code:     "/**\n * Definition for a binary tree node.\n * class TreeNode {\n *     val: number\n *     left: TreeNode | null\n *     right: TreeNode | null\n *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {\n *         this.val = (val===undefined ? 0 : val)\n *         this.left = (left===undefined ? null : left)\n *         this.right = (right===undefined ? null : right)\n *     }\n * }\n */\n\nfunction sortedArrayToBST(nums: number[]): TreeNode | null {\n\n};",
		},
		{
			Lang:     "Racket",
			LangSlug: "racket",
			Code:     "; Definition for a binary tree node.\n#|\n\n; val : integer?\n; left : (or/c tree-node? #f)\n; right : (or/c tree-node? #f)\n(struct tree-node\n  (val left right) #:mutable #:transparent)\n\n; constructor\n(define (make-tree-node [val 0])\n  (tree-node val #f #f))\n\n|#\n\n(define/contract (sorted-array-to-bst nums)\n  (-> (listof exact-integer?) (or/c tree-node? #f))\n\n  )",
		},
		{
			Lang:     "Erlang",
			LangSlug: "erlang",
			Code:     "%% Definition for a binary tree node.\n%%\n%% -record(tree_node, {val = 0 :: integer(),\n%%                     left = null  :: 'null' | #tree_node{},\n%%                     right = null :: 'null' | #tree_node{}}).\n\n-spec sorted_array_to_bst(Nums :: [integer()]) -> #tree_node{} | null.\nsorted_array_to_bst(Nums) ->\n  .",
		},
		{
			Lang:     "Elixir",
			LangSlug: "elixir",
			Code:     "# Definition for a binary tree node.\n#\n# defmodule TreeNode do\n#   @type t :: %__MODULE__{\n#           val: integer,\n#           left: TreeNode.t() | nil,\n#           right: TreeNode.t() | nil\n#         }\n#   defstruct val: 0, left: nil, right: nil\n# end\n\ndefmodule Solution do\n  @spec sorted_array_to_bst(nums :: [integer]) :: TreeNode.t | nil\n  def sorted_array_to_bst(nums) do\n\n  end\nend",
		},
	},
}
