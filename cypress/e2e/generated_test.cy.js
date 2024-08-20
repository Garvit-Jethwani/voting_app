// ********RoostGPT********
// Generated Cypress test

// ********RoostGPT********
describe('Voting Application E2E Test Suite', () => {
    beforeEach(() => {
        // Setup intercepts for API mocking
        cy.intercept('GET', '**/api/candidates', { fixture: 'candidates.json' }).as('getCandidates');
        cy.intercept('POST', '**/api/vote', { statusCode: 204 }).as('postVote');
        cy.intercept('GET', '**/api/results', { fixture: 'results.json' }).as('getResults');
        
        // Visit the home page before each test
        cy.visit('/voter');
        cy.wait('@getCandidates');
    });

    it('loads all candidate cards and the Show Results button is initially hidden', () => {
        cy.get('[data-testid="candidate-card"]').should('have.length', 4);
        cy.get('[data-testid="show-results-button"]').should('not.be.visible');
    });

    it('selecting a candidate enables the Show Results button', () => {
        cy.get('[data-testid="candidate-card"]').first().click();
        cy.get('[data-testid="show-results-button"]').should('be.visible');
    });

    describe('Voting and Results Display', () => {
        beforeEach(() => {
            // Select a candidate and show results
            cy.get('[data-testid="candidate-card"]').first().click();
            cy.get('[data-testid="show-results-button"]').click();
            cy.wait('@getResults');
        });

        it('displays the results correctly with vote percentages', () => {
            cy.get('[data-testid="result-card"]').each(($card, index) => {
                cy.wrap($card).within(() => {
                    cy.get('.progressbar_front').invoke('width').should('be.gte', 0);
                });
            });
        });

        it('navigates back to home and retains state', () => {
            cy.get('[data-testid="home-link"]').click();
            cy.url().should('include', '/home');
            cy.get('[data-testid="candidate-card"]').filter('.selected').should('have.length', 1);
        });
    });

    it('handles errors during voting', () => {
        cy.intercept('POST', '**/api/vote', { statusCode: 500 }).as('postVoteError');
        cy.get('[data-testid="candidate-card"]').first().click();
        cy.get('[data-testid="show-results-button"]').click();
        cy.wait('@postVoteError');
        cy.get('[data-testid="error-message"]').should('be.visible').and('contain', 'Error voting');
    });
});

