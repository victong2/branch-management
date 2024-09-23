# Take-Home Exercise: Branch Management API in Go

## Important

- Please try to keep time spent under **2 hours** to complete this exercise.
- 100% completion is not important, this exercise is meant to evaluate how you approach the problem and the decisions you make.

## Objective

Create a RESTful API in Go that manages the structure of a business with multiple branches. Each branch has specific requirements and restrictions for employees to be allowed to work there. These requirements and restrictions are inherited from parent branches up the hierarchy.

## Requirements

- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Time Limit**: 2 hours (timeboxed)

### Functional Requirements

1. **Branch Creation**

   - Implement an endpoint to create new branches.
   - Each branch should have:
     - A unique identifier.
     - A name.
     - An optional parent branch (null for the root branch).
     - A list of requirements and restrictions (e.g., certifications, training).

2. **Branch Hierarchy**

   - The company structure is a non-cyclical hierarchy (tree structure).
   - Each branch has only one parent but can have multiple child branches.
   - Requirements and restrictions are inherited from all ancestor branches, parents, grandparents etc.

3. **Requirements Query**
   - Implement an endpoint to retrieve the cumulative list of requirements and restrictions for a specific branch.
   - The list should include the branch's own requirements plus all inherited requirements from its ancestor branches.

## Submission

- Upload your code to a public Git repository (GitHub, GitLab, etc.).
- Send us the link to the repository so that we book some time to go over it together.

## Tips

- Focus on satisfying the requirements before optimizing or working on extra features.
- Commit small pieces often.
- Make sure you can explain your decisions and tradeoffs that were made.
