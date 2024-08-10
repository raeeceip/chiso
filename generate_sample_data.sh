#!/bin/bash

# Create directory structure
mkdir -p personal_data/{journals,blog_posts,social_media,references}
mkdir -p beliefs

# Function to generate random tags
generate_tags() {
    tags=("#AI" "#personalGrowth" "#technology" "#learning" "#productivity" "#health" "#creativity" "#mindfulness")
    num_tags=$((RANDOM % 3 + 1))
    selected_tags=""
    for i in $(seq 1 $num_tags); do
        random_tag=${tags[$RANDOM % ${#tags[@]}]}
        selected_tags="$selected_tags $random_tag"
    done
    echo $selected_tags
}

# Generate journal entries
for i in {1..10}; do
    date=$(date -d "${i} days ago" +%Y-%m-%d)
    tags=$(generate_tags)
    echo "Dear Diary,

Today I reflected on my personal growth journey. $tags

I've been thinking about how I can improve my skills and become more productive. 
There's so much to learn, especially in the field of AI and technology.

I believe that continuous learning is the key to success in any field.
I believe in the power of collaboration and sharing knowledge with others.

I'm excited about the possibilities and looking forward to what tomorrow brings." > personal_data/journals/${date}_daily_thoughts.txt
done

# Generate blog posts
for i in {1..5}; do
    date=$(date -d "${i} months ago" +%Y-%m-%d)
    tags=$(generate_tags)
    echo "The Future of AI: Possibilities and Challenges

$tags

Artificial Intelligence is rapidly evolving, opening up new possibilities across various industries. 
From healthcare to finance, AI is transforming the way we work and live.

However, with great power comes great responsibility. We must consider the ethical implications 
of AI and ensure that it's developed and used in ways that benefit humanity as a whole.

What are your thoughts on the future of AI?" > personal_data/blog_posts/${date}_ai_thoughts.txt
done

# Generate social media posts
for i in {1..15}; do
    date=$(date -d "${i} days ago" +%Y-%m-%d)
    tags=$(generate_tags)
    echo "Just finished reading a fascinating article on machine learning applications in healthcare. 
The potential for AI to improve patient outcomes is truly exciting! $tags 
#thoughtLeadership" > personal_data/social_media/${date}_twitter_post.txt
done

# Generate references
for i in {1..7}; do
    date=$(date -d "${i} months ago" +%Y-%m-%d)
    tags=$(generate_tags)
    echo "Title: 'The Singularity Is Near' by Ray Kurzweil

$tags

Key takeaways:
1. The pace of technological progress is exponential, not linear.
2. AI will surpass human intelligence in the near future.
3. The fusion of human and machine intelligence will lead to unprecedented advancements.

Thoughts: Kurzweil's predictions are bold and thought-provoking. While some may seem far-fetched, 
his track record of accurate predictions makes this book a must-read for anyone interested in 
the future of technology and humanity." > personal_data/references/${date}_book_notes.txt
done

# Generate belief files
generate_belief_file() {
    category=$1
    file_name="beliefs/${category// /_}.txt"
    echo "Generating beliefs for category: $category"
    case $category in
        "Fundamental Principles")
            echo "Fundamental Principles: Ethical conduct is the foundation of all actions" >> $file_name
            echo "Fundamental Principles: Continuous learning is essential for personal and societal growth" >> $file_name
            ;;
        "Knowledge and Truth")
            echo "Knowledge and Truth: The scientific method is the most reliable path to understanding" >> $file_name
            echo "Knowledge and Truth: Critical thinking should be applied to all information" >> $file_name
            ;;
        "Personal Growth")
            echo "Personal Growth: Challenges are opportunities for learning and improvement" >> $file_name
            echo "Personal Growth: Self-reflection is key to personal development" >> $file_name
            ;;
        "Interpersonal Relations")
            echo "Interpersonal Relations: Empathy is crucial for meaningful connections" >> $file_name
            echo "Interpersonal Relations: Clear communication prevents misunderstandings" >> $file_name
            ;;
        "Societal Responsibility")
            echo "Societal Responsibility: Everyone has a role in creating a better society" >> $file_name
            echo "Societal Responsibility: Environmental stewardship is a collective duty" >> $file_name
            ;;
        "Personal Values")
            echo "Personal Values: Integrity should guide all decisions and actions" >> $file_name
            echo "Personal Values: Authenticity leads to genuine happiness and fulfillment" >> $file_name
            ;;
        "Decision Making")
            echo "Decision Making: Consider long-term consequences in all choices" >> $file_name
            echo "Decision Making: Gather diverse perspectives before making important decisions" >> $file_name
            ;;
    esac
}

belief_categories=("Fundamental Principles" "Knowledge and Truth" "Personal Growth" "Interpersonal Relations" "Societal Responsibility" "Personal Values" "Decision Making")

for category in "${belief_categories[@]}"; do
    generate_belief_file "$category"
done

echo "Sample data and belief files generation complete."