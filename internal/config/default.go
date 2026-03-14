package config

const DefaultConfigTemplate = `# Commit Improver CLI configuration
#
# Required fields must be filled for the CLI to work.

# LLM provider used to improve commits. Options: 'gemini'. Default: 'gemini'
provider: gemini

# Model name used by the provider. Default: 'gemini-2.5-flash'
model: gemini-2.5-flash

# Language used for generated commits. Options: 'en' (English), 'pt-BR' (Português do Brasil) or 'es' (Español). Default: 'en'
language: en

# Maximum number of diff lines sent to the LLM. Default: 200
diff_limit: 200

# Allow the user to edit the final commit message. Default: false
allow_final_edit: false

gemini:
  # Gemini API key
  api_key: # REQUIRED (not indicated for --repo configurations, as it will be versioned with the code)

`
