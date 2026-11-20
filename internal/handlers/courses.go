package handlers

import (
	"encoding/json"
	"net/http"
)

type Image struct {
	URL string `json:"url"`
	Alt string `json:"alt"`
}

type Item struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       Image  `json:"image"`
}

// gu

type CoursesData struct {
	Development []Item `json:"Development"`
	Business    []Item `json:"Business"`
	Design      []Item `json:"Design"`
	Lifestyle   []Item `json:"Lifestyle"`
}

type CoursesResponse struct {
	Data CoursesData `json:"data"`
}

func Courses(w http.ResponseWriter, r *http.Request) {

	response := CoursesResponse{
		Data: CoursesData{

			Development: []Item{
				{
					ID:          "WD101",
					Title:       "Web Development Fundamentals",
					Description: "This course covers the basic concepts and tools for building static and dynamic websites. Students will learn HTML, CSS, JavaScript, and jQuery to create engaging and interactive web pages.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Development/Web%20Development%20Fundamentals.png",
						Alt: "Web Development Fundamentals",
					},
				},
				{
					ID:          "DS201",
					Title:       "Data Science Essentials",
					Description: "This course provides an introduction to the fundamental concepts and techniques used in data science. Students will learn how to collect, clean, analyze, and visualize data using popular tools and programming languages like Python and R.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Development/Data%20Science%20Essentials.png",
						Alt: "Data Science Essentials",
					},
				},
				{
					ID:          "ST301",
					Title:       "Software Testing Fundamentals",
					Description: "This course covers the basics of software testing, including test planning, test case design, test execution, and defect management. Students will learn how to use popular testing frameworks and tools to ensure the quality and reliability of software applications.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Development/Software%20Testing%20Fundamentals.png",
						Alt: "Software Testing Fundamentals",
					},
				},
				{
					ID:          "MD401",
					Title:       "Mobile App Development with React Native",
					Description: "This course teaches students how to build native mobile apps using the React Native framework. Students will learn how to create cross-platform apps that work on both iOS and Android devices, using a single codebase.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Development/Mobile%20App%20Development%20With%20React%20Native.png",
						Alt: "Mobile App Development with React Native",
					},
				},
				{
					ID:          "DO501",
					Title:       "DevOps for Agile Teams",
					Description: "This course provides an overview of DevOps principles and practices, with a focus on how they can be applied in agile software development teams. Students will learn how to use popular tools and techniques for continuous integration, continuous delivery, and automated testing.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Development/Devops%20for%20Agile%20Teams.png",
						Alt: "DevOps for Agile Teams",
					},
				},
			},

			Business: []Item{
				{
					ID:          "MK101",
					Title:       "Introduction to Marketing",
					Description: "This course provides an overview of marketing principles and practices, including market research, segmentation, targeting, and positioning. Students will learn how to create effective marketing plans and campaigns, using both traditional and digital marketing techniques.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Business/Introduction%20To%20Marketing.png",
						Alt: "Introduction to Marketing",
					},
				},
				{
					ID:          "LD201",
					Title:       "Leadership Development",
					Description: "This course covers the fundamental concepts and skills required for effective leadership, including communication, motivation, delegation, and team building. Students will learn how to lead teams and organizations, and how to handle common leadership challenges.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Business/Leadership%20Development.png",
						Alt: "Leadership Development",
					},
				},
				{
					ID:          "FN301",
					Title:       "Finance for Non-Financial Managers",
					Description: "This course provides an introduction to finance concepts and tools for non-financial managers. Students will learn how to read and interpret financial statements, analyze financial data, and make informed financial decisions.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Business/Finance%20for%20Non-Financial%20Managers.png",
						Alt: "Finance for Non-Financial Managers",
					},
				},
				{
					ID:          "OM401",
					Title:       "Operations Management Fundamentals",
					Description: "This course covers the basics of operations management, including process design, capacity planning, inventory management, and quality control. Students will learn how to optimize operations to improve efficiency, productivity, and customer satisfaction.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Business/Operations%20Management%20Fundamentals.png",
						Alt: "Operations Management Fundamentals",
					},
				},
				{
					ID:          "PM501",
					Title:       "Project Management Essentials",
					Description: "This course provides an overview of project management principles and techniques, including project planning, scheduling, budgeting, risk management, and stakeholder communication. Students will learn how to manage projects effectively, using both traditional and agile project management approaches.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Business/Project%20Management%20Essentials.png",
						Alt: "Project Management Essentials",
					},
				},
			},

			Design: []Item{
				{
					ID:          "GD101",
					Title:       "Graphic Design Fundamentals",
					Description: "This course covers the basic principles of graphic design, including typography, color theory, layout design, and image manipulation. Students will learn how to use industry-standard design software to create effective visual designs for print and digital media.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Design/Graphic%20Design%20Fundamentals.png",
						Alt: "Graphic Design Fundamentals",
					},
				},
				{
					ID:          "UX201",
					Title:       "User Experience Design",
					Description: "This course covers the fundamentals of user experience (UX) design, including user research, interaction design, information architecture, and usability testing. Students will learn how to design digital products and services that meet user needs and business goals.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Design/User%20Experience%20Design.png",
						Alt: "User Experience Design",
					},
				},
				{
					ID:          "ID301",
					Title:       "Industrial Design Essentials",
					Description: "This course covers the fundamental concepts and techniques used in industrial design, including sketching, prototyping, modeling, and design for manufacturing. Students will learn how to create innovative and functional products that meet user needs and market demands.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Design/Industrial%20Design%20Essentials.png",
						Alt: "Industrial Design Essentials",
					},
				},
				{
					ID:          "FD401",
					Title:       "Fashion Design Principles",
					Description: "This course covers the principles and practices of fashion design, including fashion illustration, textile design, pattern making, and garment construction. Students will learn how to design and produce original and stylish clothing for different markets and occasions.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Design/Fashion%20Design%20Principles.png",
						Alt: "Fashion Design Principles",
					},
				},
			},

			Lifestyle: []Item{
				{
					ID:          "LC101",
					Title:       "Cooking Basics",
					Description: "This course provides an introduction to cooking techniques, ingredients, and flavors. Students will learn how to plan and prepare meals, using simple recipes and kitchen tools. The course covers a range of cuisines and dietary preferences.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Lifestyle/Cooking%20Basics.png",
						Alt: "Cooking Basics",
					},
				},
				{
					ID:          "LF201",
					Title:       "Fitness for Life",
					Description: "This course covers the fundamentals of fitness, including exercise principles, workout planning, and nutrition. Students will learn how to build strength, endurance, and flexibility, and how to create a healthy and sustainable lifestyle.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Lifestyle/Fitness%20For%20Life.png",
						Alt: "Fitness for Life",
					},
				},
				{
					ID:          "LT301",
					Title:       "Travel Photography",
					Description: "This course covers the fundamentals of travel photography, including composition, lighting, and storytelling. Students will learn how to use a range of camera equipment and editing software to capture and share their travel experiences.",
					Image: Image{
						URL: "https://codehelp-apis.vercel.app/get-top-courses/Lifestyle/Travel%20Photography.png",
						Alt: "Travel Photography",
					},
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
